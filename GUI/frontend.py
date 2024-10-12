import tkinter as tk
from tkinter import ttk, scrolledtext
import requests

# Function stubs (Replace with actual API calls)
def add_book():
    book_data = {
        "title": entry_book_title.get(),
        "year": int(entry_book_year.get()),
        "author_id": int(entry_book_author_id.get()),
        "isbn": entry_book_isbn.get(),
        "publication": entry_book_publication.get(),
        "num_of_copies": int(entry_book_num_of_copies.get())
    }
    response = requests.post('http://localhost:9010/book/', json=book_data)
    text_books.insert(tk.END, f"Book added: {response.json()}\n")

def get_book_by_id():
    book_id = entry_get_book_id.get()
    response = requests.get(f'http://localhost:9010/book/{book_id}')
    book = response.json()
    text_books.delete('1.0', tk.END)  # Clear previous content
    
    formatted_book = (
            f"Title: {book['title']}\n"
            f"Year: {book['year']}\n"
            f"Author ID: {book['author_id']}\n"
            f"ISBN: {book['isbn']}\n"
            f"Publication: {book['publication']}\n"
            f"Number of Copies: {book['num_of_copies']}\n"
            "----------------------\n"
        )
    text_books.insert(tk.END, formatted_book)

def get_all_books():
    response = requests.get('http://localhost:9010/book/')
    books = response.json()
    text_books.delete('1.0', tk.END)  # Clear previous content
    for book in books:
        formatted_book = (
            #f"ID: {book['id']}\n"
            f"Title: {book['title']}\n"
            f"Year: {book['year']}\n"
            f"Author ID: {book['author_id']}\n"
            f"ISBN: {book['isbn']}\n"
            f"Publication: {book['publication']}\n"
            f"Number of Copies: {book['num_of_copies']}\n"
            "----------------------\n"
        )
        text_books.insert(tk.END, formatted_book)

def delete_book():
    book_id = entry_get_book_id.get()
    response = requests.delete(f'http://localhost:9010/book/{book_id}')
    text_books.insert(tk.END, f"Book deleted: {response.json()}\n")

def add_author():
    author_data = {
        "name": entry_author_name.get(),
        "email": entry_author_email.get()
    }
    response = requests.post('http://localhost:9010/author/', json=author_data)
    text_authors.insert(tk.END, f"Author added: {response.json()}\n")

def get_author_by_id():
    author_id = entry_get_author_id.get()
    response = requests.get(f'http://localhost:9010/author/{author_id}')
    text_authors.insert(tk.END, f"Author: {response.json()}\n")

def get_all_authors():
    response = requests.get('http://localhost:9010/author/')
    authors = response.json()
    text_authors.delete('1.0', tk.END)  # Clear previous content
    for author in authors:
        formatted_author = (
            #f"ID: {author['id']}\n"
            f"Name: {author['name']}\n"
            f"Email: {author['email']}\n"
            "----------------------\n"
        )
        text_authors.insert(tk.END, formatted_author)

def delete_author():
    author_id = entry_get_author_id.get()
    response = requests.delete(f'http://localhost:9010/author/{author_id}')
    text_authors.insert(tk.END, f"Author deleted: {response.json()}\n")

def add_magazine():
    magazine_data = {
        "title": entry_magazine_title.get(),
        "issue_number": int(entry_magazine_issue.get()),
        "num_of_copies": int(entry_magazine_copies.get()),
        "publisher": entry_magazine_publisher.get(),
        "year": int(entry_magazine_year.get())
    }
    response = requests.post('http://localhost:9010/magazine/', json=magazine_data)
    text_magazines.insert(tk.END, f"Magazine added: {response.json()}\n")

def get_magazine_by_id():
    magazine_id = entry_get_magazine_id.get()
    response = requests.get(f'http://localhost:9010/magazine/{magazine_id}')
    magazine = response.json()
    text_magazines.delete('1.0', tk.END)  # Clear previous content
    formatted_magazine = (
            #f"ID: {magazine['id']}\n"
            f"Title: {magazine['title']}\n"
            f"Issue Number: {magazine['issue_number']}\n"
            f"Number of Copies: {magazine['num_of_copies']}\n"
            f"Publisher: {magazine['publisher']}\n"
            f"Year: {magazine['year']}\n"
            "----------------------\n"
        )
    text_magazines.insert(tk.END, f"Magazine: {response.json()}\n")

def get_all_magazines():
    response = requests.get('http://localhost:9010/magazine/')
    magazines = response.json()
    text_magazines.delete('1.0', tk.END)  # Clear previous content
    for magazine in magazines:
        formatted_magazine = (
            #f"ID: {magazine['id']}\n"
            f"Title: {magazine['title']}\n"
            f"Issue Number: {magazine['issue_number']}\n"
            f"Number of Copies: {magazine['num_of_copies']}\n"
            f"Publisher: {magazine['publisher']}\n"
            f"Year: {magazine['year']}\n"
            "----------------------\n"
        )
        text_magazines.insert(tk.END, formatted_magazine)

def delete_magazine():
    magazine_id = entry_get_magazine_id.get()
    response = requests.delete(f'http://localhost:9010/magazine/{magazine_id}')
    text_magazines.insert(tk.END, f"Magazine deleted: {response.json()}\n")

def create_member():
    member_data = {
        "name": entry_member_name.get(),
        "email": entry_member_email.get(),
        "outdated_fees": float(entry_member_outdatedFees.get())
    }
    response = requests.post('http://localhost:9010/member/', json=member_data)
    text_members.insert(tk.END, f"Member created: {response.json()}\n")

def get_member_fees():
    member_id = entry_get_member_id.get()
    response = requests.get(f'http://localhost:9010/members/{member_id}/fees')
    
    if response.status_code == 200:  # Check if the request was successful
        member_data = response.json()
        text_members.delete('1.0', tk.END)  # Clear previous content
        
        # Extract member fees and format them
        member = member_data["member"]
        formatted_member = (
            f"Name: {member['name']}\n"
            f"Email: {member['email']}\n"
            f"Outdated Fees: {member_data['outdated_fees']}\n"
            f"Overdue Days: {member_data['overdue_days']}\n"
            f"Total Amount Due: {member_data['total_amount']}\n"
            "----------------------\n"
        )
        text_members.insert(tk.END, formatted_member)
    else:
        text_members.insert(tk.END, f"Error: Could not retrieve fees for Member ID {member_id}\n")

def get_loans_by_member_id():
    member_id = entry_get_member_id.get()
    response = requests.get(f'http://localhost:9010/member/{member_id}/loans')
    text_loans.insert(tk.END, f"Loans for Member {member_id}: {response.json()}\n")

def create_loan():
    loan_data = {
        "member_id": int(entry_loan_member_id.get()),
        "loanable_id": int(entry_loan_loanable_id.get()),
        "loanable_type": entry_loan_type.get(),
    }
    response = requests.post('http://localhost:9010/loan/', json=loan_data)
    text_loans.insert(tk.END, f"Loan created: {response.json()}\n")
'''
def get_all_loans():
    response = requests.get('http://localhost:9010/loan/')
    loans = response.json()
    text_loans.delete('1.0', tk.END)  # Clear previous content
    for loan in loans:
        formatted_loan = (
            f"Loan ID: {loan['id']}\n"
            f"Member ID: {loan['member_id']}\n"
            f"Loanable ID: {loan['loanable_id']}\n"
            f"Loanable Type: {loan['loanable_type']}\n"
            f"Borrow Date: {loan['borrow_date']}\n"
            f"Return Date: {loan['return_date']}\n"
            "----------------------\n"
        )
        text_loans.insert(tk.END, formatted_loan)
'''
# Main application window
app = tk.Tk()
app.title("Library Management System")

notebook = ttk.Notebook(app)
notebook.pack(expand=True, fill='both')

# Book Tab
tab_books = ttk.Frame(notebook)
notebook.add(tab_books, text="Books")

tk.Label(tab_books, text="Title:").grid(row=0, column=0)
entry_book_title = tk.Entry(tab_books)
entry_book_title.grid(row=0, column=1)

tk.Label(tab_books, text="Year:").grid(row=1, column=0)
entry_book_year = tk.Entry(tab_books)
entry_book_year.grid(row=1, column=1)

tk.Label(tab_books, text="Author ID:").grid(row=2, column=0)
entry_book_author_id = tk.Entry(tab_books)
entry_book_author_id.grid(row=2, column=1)

tk.Label(tab_books, text="ISBN:").grid(row=3, column=0)
entry_book_isbn = tk.Entry(tab_books)
entry_book_isbn.grid(row=3, column=1)

tk.Label(tab_books, text="Publication:").grid(row=4, column=0)
entry_book_publication = tk.Entry(tab_books)
entry_book_publication.grid(row=4, column=1)

tk.Label(tab_books, text="Num of Copies:").grid(row=5, column=0)
entry_book_num_of_copies = tk.Entry(tab_books)
entry_book_num_of_copies.grid(row=5, column=1)

tk.Button(tab_books, text="Add Book", command=add_book).grid(row=6, column=0)
tk.Button(tab_books, text="Get All Books", command=get_all_books).grid(row=6, column=1)

# Book ID for search/delete
tk.Label(tab_books, text="Book ID:").grid(row=7, column=0)
entry_get_book_id = tk.Entry(tab_books)
entry_get_book_id.grid(row=7, column=1)
tk.Button(tab_books, text="Get Book", command=get_book_by_id).grid(row=8, column=0)
tk.Button(tab_books, text="Delete Book", command=delete_book).grid(row=8, column=1)

text_books = scrolledtext.ScrolledText(tab_books, width=50, height=10)
text_books.grid(row=9, column=0, columnspan=2)

# Author Tab
tab_authors = ttk.Frame(notebook)
notebook.add(tab_authors, text="Authors")

tk.Label(tab_authors, text="Name:").grid(row=0, column=0)
entry_author_name = tk.Entry(tab_authors)
entry_author_name.grid(row=0, column=1)

tk.Label(tab_authors, text="Email:").grid(row=1, column=0)
entry_author_email = tk.Entry(tab_authors)
entry_author_email.grid(row=1, column=1)

tk.Button(tab_authors, text="Add Author", command=add_author).grid(row=2, column=0)
tk.Button(tab_authors, text="Get All Authors", command=get_all_authors).grid(row=2, column=1)

# Author ID for search/delete
tk.Label(tab_authors, text="Author ID:").grid(row=3, column=0)
entry_get_author_id = tk.Entry(tab_authors)
entry_get_author_id.grid(row=3, column=1)
tk.Button(tab_authors, text="Get Author", command=get_author_by_id).grid(row=4, column=0)
tk.Button(tab_authors, text="Delete Author", command=delete_author).grid(row=4, column=1)

text_authors = scrolledtext.ScrolledText(tab_authors, width=50, height=10)
text_authors.grid(row=5, column=0, columnspan=2)

# Magazine Tab
tab_magazines = ttk.Frame(notebook)
notebook.add(tab_magazines, text="Magazines")

tk.Label(tab_magazines, text="Title:").grid(row=0, column=0)
entry_magazine_title = tk.Entry(tab_magazines)
entry_magazine_title.grid(row=0, column=1)

tk.Label(tab_magazines, text="Issue Number:").grid(row=1, column=0)
entry_magazine_issue = tk.Entry(tab_magazines)
entry_magazine_issue.grid(row=1, column=1)

tk.Label(tab_magazines, text="Num of Copies:").grid(row=2, column=0)
entry_magazine_copies = tk.Entry(tab_magazines)
entry_magazine_copies.grid(row=2, column=1)

tk.Label(tab_magazines, text="Publisher:").grid(row=3, column=0)
entry_magazine_publisher = tk.Entry(tab_magazines)
entry_magazine_publisher.grid(row=3, column=1)

tk.Label(tab_magazines, text="Year:").grid(row=4, column=0)
entry_magazine_year = tk.Entry(tab_magazines)
entry_magazine_year.grid(row=4, column=1)

tk.Button(tab_magazines, text="Add Magazine", command=add_magazine).grid(row=5, column=0)
tk.Button(tab_magazines, text="Get All Magazines", command=get_all_magazines).grid(row=5, column=1)

# Magazine ID for search/delete
tk.Label(tab_magazines, text="Magazine ID:").grid(row=6, column=0)
entry_get_magazine_id = tk.Entry(tab_magazines)
entry_get_magazine_id.grid(row=6, column=1)
tk.Button(tab_magazines, text="Get Magazine", command=get_magazine_by_id).grid(row=7, column=0)
tk.Button(tab_magazines, text="Delete Magazine", command=delete_magazine).grid(row=7, column=1)

text_magazines = scrolledtext.ScrolledText(tab_magazines, width=50, height=10)
text_magazines.grid(row=8, column=0, columnspan=2)

# Member Tab
tab_members = ttk.Frame(notebook)
notebook.add(tab_members, text="Members")

tk.Label(tab_members, text="Name:").grid(row=0, column=0)
entry_member_name = tk.Entry(tab_members)
entry_member_name.grid(row=0, column=1)

tk.Label(tab_members, text="Email:").grid(row=1, column=0)
entry_member_email = tk.Entry(tab_members)
entry_member_email.grid(row=1, column=1)

tk.Label(tab_members, text="Outdated Fees:").grid(row=2, column=0)
entry_member_outdatedFees = tk.Entry(tab_members)
entry_member_outdatedFees.grid(row=2, column=1)

tk.Button(tab_members, text="Create Member", command=create_member).grid(row=3, column=0)

# Member ID for searching loans
tk.Label(tab_members, text="Member ID:").grid(row=4, column=0)
entry_get_member_id = tk.Entry(tab_members)
entry_get_member_id.grid(row=4, column=1)
tk.Button(tab_members, text="Get Member Loans", command=get_loans_by_member_id).grid(row=5, column=0)

# Button to fetch member fees
button_get_member_fees = ttk.Button(tab_members, text="Get Member Fees", command=get_member_fees)
button_get_member_fees.grid(row=5, column=1, columnspan=2)

text_members = scrolledtext.ScrolledText(tab_members, width=50, height=10)
text_members.grid(row=6, column=0, columnspan=2)

# Loan Tab
tab_loans = ttk.Frame(notebook)
notebook.add(tab_loans, text="Loans")

tk.Label(tab_loans, text="Member ID:").grid(row=0, column=0)
entry_loan_member_id = tk.Entry(tab_loans)
entry_loan_member_id.grid(row=0, column=1)

tk.Label(tab_loans, text="Loanable ID:").grid(row=1, column=0)
entry_loan_loanable_id = tk.Entry(tab_loans)
entry_loan_loanable_id.grid(row=1, column=1)

tk.Label(tab_loans, text="Type (book/magazine):").grid(row=2, column=0)
entry_loan_type = tk.Entry(tab_loans)
entry_loan_type.grid(row=2, column=1)

tk.Button(tab_loans, text="Create Loan", command=create_loan).grid(row=3, column=0)
#tk.Button(tab_loans, text="Get All Loans", command=get_all_loans).grid(row=3, column=1)

text_loans = scrolledtext.ScrolledText(tab_loans, width=50, height=10)
text_loans.grid(row=4, column=0, columnspan=2)

app.mainloop()
