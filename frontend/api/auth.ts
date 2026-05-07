import { BASE_URL } from "@/lib/constants";
import axios from "axios";
export const registerRequest = async (data: any) => {
  const response = await axios.post(`${BASE_URL}api/v1/signup`, { 
  
      name: data.name,
      email: data.email,
      password: data.password,
  });
  console.log(data);
  return response.data;
};

export const loginRequest = async (data: any) => {
 const response = await axios.post(`${BASE_URL}api/v1/login`, data);
  return response.data;
};
