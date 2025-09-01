import React from "react";
import Navbar from "../../functions/navbar";
import "../../css/cart.css";
import { Breadcrumb, Input, Checkbox, Modal } from "antd";
import { HomeOutlined, ShoppingOutlined } from "@ant-design/icons";

const Cart = () => {
  return (
    <div>
      <Navbar />
      <div className="background-cart">
        <Breadcrumb className=".font-cart"
          items={[
            { href: "", title: <HomeOutlined /> },
            {
              title: (
                <>
                  <ShoppingOutlined /> Cart
                </>
              ),
            },
          ]}
        />
      <div className="search-container">
          <Input.Search
            className="custom-placeholder"
            placeholder="Search for products,brands and shops"
            allowClear
            onSearch={(value) => console.log("Search:", value)}
            style={{ width: 500 }}
          />
        </div>
      </div>
    </div>
  );
};

export default Cart;
