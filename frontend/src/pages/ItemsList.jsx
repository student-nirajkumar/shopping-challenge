import { useEffect, useState } from "react";
import axios from "../api/index";
import NavBar from "../components/NavBar";

export default function ItemsList() {
  const [items, setItems] = useState([]);

  // Fetch items on component mount
  useEffect(() => {
    axios
      .get("/items")
      .then((res) => {
        setItems(res.data);
      })
      .catch(() => {
        alert("Failed to load items");
      });
  }, []);

  // Add item to cart
  function addToCart(itemId) {
    axios
      .post("/carts", { item_id: itemId })
      .then(() => {
        alert("Item added to cart!");
      })
      .catch(() => {
        alert("Failed to add item");
      });
  }

  return (
    <div>
      <NavBar />
      <h2 style={{ textAlign: "center" }}>Items List</h2>

      <div style={{ padding: "20px" }}>
        {items.length === 0 && <p>No items found.</p>}

        {items.map((item) => (
          <div
            key={item.ID}
            style={{
              padding: "10px",
              marginBottom: "10px",
              border: "1px solid gray",
              borderRadius: "6px",
            }}
          >
            <strong>{item.Name}</strong> — {item.Status}
            <button
              style={{ marginLeft: "10px" }}
              onClick={() => addToCart(item.ID)}
            >
              Add to Cart
            </button>
          </div>
        ))}
      </div>
    </div>
  );
}
