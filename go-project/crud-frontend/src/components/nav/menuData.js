// menuData.js
import { FaTasks, FaPlus, FaUserPlus } from "react-icons/fa"; // Import icons from react-icons or your preferred library

const menuData = [
  {
    name: "Task List",
    href: "/",
    icon: <FaTasks />,
  },
  {
    name: "Create Task",
    href: "/create",
    icon: <FaPlus />,
  },
  {
    name: "Sign Up",
    href: "/users/signup",
    icon: <FaUserPlus />,
  },
];

export default menuData;
