import React from "react";
import "../css/navbar.css"

const Navbar = () => {
    return(
        <div>
            <div className="background-navbar">
                <div className="size-header">
                    <div className="font-navbar">Seller</div>
                    <div>|</div>
                    <div>Member</div>
                </div>
            </div>
        </div>
    );
};
export default Navbar;