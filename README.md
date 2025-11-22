🚀 How to Run the Backend
1. Go to the backend folder
cd backend

2. Install Go dependencies
go mod tidy

3. Make sure PostgreSQL is running

Database details used:

Setting	Value
host	localhost
port	5432
user	postgres
password	root
database	shopping

Create DB if not exists:

CREATE DATABASE shopping;

4. Run backend server
go run ./cmd


Server will run at:

👉 http://localhost:8080

🔑 API Endpoints Summary
✔ Auth
Method	Endpoint	Description
POST	/users	Signup
POST	/users/login	Login (returns token)
✔ Items
Method	Endpoint	Description
POST	/items	Create item
GET	/items	List items
✔ Cart

All require Authorization: Bearer <token>

Method	Endpoint	Description
POST	/carts	Add item to cart
GET	/carts/me	View own cart
GET	/carts	List all carts
✔ Orders
Method	Endpoint	Description
POST	/orders	Create order
GET	/orders/me	List my orders
GET	/orders	List all orders
🌐 How to Run the Frontend (React)
1. Go to frontend folder
cd frontend

2. Install dependencies
npm install

3. Start the dev server
npm run dev


Frontend will run at:

👉 http://localhost:5173

🔗 Environment Variables (Frontend)

In src/api/index.js, API base URL is hardcoded:

const api = axios.create({
  baseURL: "http://localhost:8080",
});

💾 Database Structure Summary
Users
id, username, password_hash, token, created_at

Items
id, name, status, created_at

Carts
id, user_id, status, created_at

CartItems
cart_id, item_id

Orders
id, cart_id, user_id, created_at

🧪 Testing (Thunder Client)

Endpoints tested using:

Signup

Login

Create Item

Add to Cart

Checkout

View Orders