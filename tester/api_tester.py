#!/usr/bin/env python3
import requests
import json
import sys
from typing import Optional, Dict, Any

class APITester:
    def __init__(self, base_url: str = "http://localhost:3000"):
        self.base_url = base_url
        self.token: Optional[str] = None
        self.username: Optional[str] = None
        self.is_admin: bool = False
        
    def make_request(self, method: str, endpoint: str, data: Optional[Dict] = None) -> Dict[str, Any]:
        url = f"{self.base_url}{endpoint}"
        headers = {}
        
        if self.token:
            headers["Authorization"] = f"Bearer {self.token}"
        if data:
            headers["Content-Type"] = "application/json"
            
        try:
            if method.upper() == "GET":
                response = requests.get(url, headers=headers)
            elif method.upper() == "POST":
                response = requests.post(url, headers=headers, json=data)
            elif method.upper() == "DELETE":
                response = requests.delete(url, headers=headers)
            else:
                return {"error": f"Unsupported method: {method}"}
                
            return {
                "status_code": response.status_code,
                "data": response.json() if response.content else {}
            }
        except requests.exceptions.ConnectionError:
            return {"error": "Cannot connect to API server"}
        except requests.exceptions.RequestException as e:
            return {"error": f"Request error: {str(e)}"}
        except json.JSONDecodeError:
            return {"error": "Invalid JSON response", "status_code": response.status_code}

    def register(self, username: str, password: str) -> bool:
        data = {"username": username, "password": password}
        response = self.make_request("POST", "/api/user/register", data)
        
        if "error" in response:
            print(f"❌ Error: {response['error']}")
            return False
            
        if response["status_code"] == 200:
            self.token = response["data"].get("data", {}).get("token")
            self.username = username
            print(f"✅ Registered! Token: {self.token[:20]}...")
            self.is_admin = True
            return True
        else:
            message = response.get("data", {}).get("message", "Registration failed")
            print(f"❌ {message}")
            return False
    
    def login(self, username: str, password: str) -> bool:
        data = {"username": username, "password": password}
        response = self.make_request("POST", "/api/user/login", data)
        
        if "error" in response:
            print(f"❌ Error: {response['error']}")
            return False
            
        if response["status_code"] == 200:
            self.token = response["data"].get("data", {}).get("token")
            self.username = username
            print(f"✅ Logged in! Token: {self.token[:20]}...")
            return True
        else:
            message = response.get("data", {}).get("message", "Login failed")
            print(f"❌ {message}")
            return False
    
    def delete_user(self) -> bool:
        if not self.token:
            print("❌ Login first!")
            return False
            
        response = self.make_request("DELETE", "/api/user/delete")
        
        if "error" in response:
            print(f"❌ Error: {response['error']}")
            return False
            
        if response["status_code"] == 200:
            print("✅ Account deleted!")
            self.token = None
            self.username = None
            self.is_admin = False
            return True
        else:
            message = response.get("data", {}).get("message", "Deletion failed")
            print(f"❌ {message}")
            return False

    def get_user_info(self) -> bool:
        if not self.token:
            print("❌ Login first!")
            return False
            
        response = self.make_request("GET", "/api/user/info")
        
        if "error" in response:
            print(f"❌ Error: {response['error']}")
            return False
            
        if response["status_code"] == 200:
            user_data = response["data"].get("data", {})
            role = user_data.get('role', 1)
            self.is_admin = (role == 0)
            
            print(f"\n👤 {user_data.get('username', 'N/A')} ({'Admin' if self.is_admin else 'Member'})")
            print(f"   Borrowed: {len(user_data.get('books_taken', []))}")
            return True
        else:
            message = response.get("data", {}).get("message", "Failed to get user info")
            print(f"❌ {message}")
            return False

    def get_all_books(self) -> bool:
        if not self.token:
            print("❌ Login first!")
            return False
            
        response = self.make_request("GET", "/api/books")
        
        if "error" in response:
            print(f"❌ Error: {response['error']}")
            return False
            
        if response["status_code"] == 200:
            books = response["data"].get("data", [])
            if not books:
                print("📚 No books available")
                return True
                
            print(f"\n📚 Books ({len(books)}):")
            for i, book in enumerate(books, 1):
                available = book.get('count', 0) - book.get('borrowed', 0)
                print(f"{i:2d}. {book.get('title', 'N/A')} - {book.get('author', 'N/A')} ({available} available)")
            return True
        else:
            message = response.get("data", {}).get("message", "Failed to get books")
            print(f"❌ {message}")
            return False
    
    def add_book(self, title: str, author: str) -> bool:
        if not self.token:
            print("❌ Login first!")
            return False
            
        if not self.is_admin:
            print("❌ Admin required!")
            return False
            
        data = {"title": title, "author": author}
        response = self.make_request("POST", "/api/book/add", data)
        
        if "error" in response:
            print(f"❌ Error: {response['error']}")
            return False
            
        if response["status_code"] == 200:
            print(f"✅ Added: {title} - {author}")
            return True
        else:
            message = response.get("data", {}).get("message", "Failed to add book")
            print(f"❌ {message}")
            return False

    def borrow_book(self, title: str) -> bool:
        if not self.token:
            print("❌ Login first!")
            return False
            
        data = {"title": title}
        response = self.make_request("POST", "/api/book/borrow", data)
        
        if "error" in response:
            print(f"❌ Error: {response['error']}")
            return False
            
        if response["status_code"] == 200:
            print(f"✅ Borrowed: {title}")
            return True
        else:
            message = response.get("data", {}).get("message", "Failed to borrow book")
            print(f"❌ {message}")
            return False
    
    def return_book(self, title: str) -> bool:
        if not self.token:
            print("❌ Login first!")
            return False
            
        data = {"title": title}
        response = self.make_request("POST", "/api/book/return", data)
        
        if "error" in response:
            print(f"❌ Error: {response['error']}")
            return False
            
        if response["status_code"] == 200:
            print(f"✅ Returned: {title}")
            return True
        else:
            message = response.get("data", {}).get("message", "Failed to return book")
            print(f"❌ {message}")
            return False

    def check_connection(self) -> bool:
        response = self.make_request("GET", "/api/books")
        return "error" not in response


def print_menu():
    print("\n📚 Library API")
    print("1. Register  2. Login  3. Delete  4. Info")
    print("5. Books     6. Add    7. Borrow  8. Return")
    print("9. Test     10. Exit")


def get_input(prompt: str) -> str:
    try:
        return input(prompt).strip()
    except KeyboardInterrupt:
        print("\n👋 Bye!")
        sys.exit(0)

def main():
    tester = APITester()
    
    print("📚 Library API Tester")
    if not tester.check_connection():
        print("❌ Server not running! Start with: go run main.go")
        return
    print("✅ Connected!")
    
    while True:
        print_menu()
        
        try:
            choice = get_input("Choice: ")
            
            if choice == "1":
                username = get_input("Username: ")
                password = get_input("Password: ")
                tester.register(username, password)
                
            elif choice == "2":
                username = get_input("Username: ")
                password = get_input("Password: ")
                tester.login(username, password)
                
            elif choice == "3":
                if tester.username:
                    confirm = get_input(f"Delete '{tester.username}'? (y/n): ")
                    if confirm.lower() in ['y', 'yes']:
                        tester.delete_user()
                else:
                    print("❌ Login first!")
                    
            elif choice == "4":
                tester.get_user_info()
                    
            elif choice == "5":
                tester.get_all_books()
                    
            elif choice == "6":
                if tester.username and tester.is_admin:
                    title = get_input("Title: ")
                    author = get_input("Author: ")
                    tester.add_book(title, author)
                else:
                    print("❌ Admin required!")
                    
            elif choice == "7":
                if tester.username:
                    title = get_input("Title: ")
                    tester.borrow_book(title)
                else:
                    print("❌ Login first!")
                    
            elif choice == "8":
                if tester.username:
                    title = get_input("Title: ")
                    tester.return_book(title)
                else:
                    print("❌ Login first!")
                    
            elif choice == "9":
                if tester.check_connection():
                    print("✅ Connected!")
                else:
                    print("❌ Disconnected!")
                    
            elif choice == "10":
                print("👋 Bye!")
                break
                
            else:
                print("❌ Invalid choice!")
                
        except Exception as e:
            print(f"❌ Error: {str(e)}")


if __name__ == "__main__":
    main()