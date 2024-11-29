import React, { useState } from "react";
import menuData from "./menuData";

export default function ResponsiveNavbar() {
  const [isSidebarOpen, setIsSidebarOpen] = useState(false);

  const toggleSidebar = () => {
    setIsSidebarOpen(!isSidebarOpen);
  };

  return (
    <div className="relative navBoxShadow">
      {/* Toggle Button for Small Screens */}
      <button
        onClick={toggleSidebar}
        className="fixed top-4 left-4 z-50 p-2 bg-black text-white rounded-md sm:hidden"
      >
        {isSidebarOpen ? "Close Menu" : "Open Menu"}
      </button>

      {/* Sidebar */}
      <div
        className={`fixed top-0 left-0 w-1/5 h-full bg-black text-white navBoxShadow flex flex-col justify-center items-center space-y-4
          ${isSidebarOpen ? "block" : "hidden"} sm:block`}
      >
        {menuData.map((item, index) => (
          <div key={index} className="flex items-center space-x-2">
            <span>{item.icon}</span>
            <a href={item.href} className="text-lg">
              {item.name}
            </a>
          </div>
        ))}
      </div>
    </div>
  );
}
