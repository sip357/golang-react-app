import axios from "axios";

const API_URL = "http://localhost:8080"; // Go backend URL

export const fetchTasks = async () => {
  const response = await axios.get(`${API_URL}/`);
  return response.data;
};

export const createTask = async (task) => {
  const response = await axios.post(`${API_URL}/create`, task);
  return response.data;
};

export const updateTask = async (task) => {
  const response = await axios.put(`${API_URL}/update`, task);
  return response.data;
};

export const deleteTask = async (id) => {
  const response = await axios.delete(`${API_URL}/delete?id=${id}`);
  return response.data;
};

export const loginHandler = async (user) => {
  const response = await axios.post(`${API_URL}/users/login`, user);
  return response.data;
};

export const signUp = async (newUser) => {
  const response = await axios.post(`${API_URL}/users/signUp`, newUser);
  return response.data;
};