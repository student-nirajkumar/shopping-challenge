import { useNavigate } from "react-router-dom";
import axios from "../api/index";

export default function NavBar() {
  const navigate = useNavigate();

  function logout() {
    localStorage.removeItem("token");
    navigate("/");
  }

  function viewCart() {
    axios
      .get("/carts/me")
      .then((res) => {
        const cart = res.data;
        let message = "Your Cart:\n";
        message += "Cart ID: " + cart.cart.ID + "\n";

        cart.items.forEach((i) => {
          message += "- Item ID: " + i.ItemID + "\n";
        });

        alert(message);
      })
      .catch(() => alert("Cannot load cart"));
  }

  function orderHistory() {
    axios
      .get("/orders/me")
      .then((res) => {
        let msg = "Your Orders:\n";
        res.data.forEach((o) => {
          msg += "- Order ID: " + o.ID + "\n";
        });
        alert(msg);
      })
      .catch(() => alert("Cannot load orders"));
  }

  function checkout() {
    axios
      .post("/orders")
      .then(() => {
        alert("Order successful!");
      })
      .catch(() => alert("Checkout failed"));
  }

  return (
    <div style={{ padding: "10px", background: "#eee" }}>
      <button onClick={viewCart}>Cart</button>
      <button onClick={orderHistory} style={{ marginLeft: "10px" }}>
        Orders
      </button>
      <button onClick={checkout} style={{ marginLeft: "10px" }}>
        Checkout
      </button>
      <button onClick={logout} style={{ marginLeft: "10px" }}>
        Logout
      </button>
    </div>
  );
}
