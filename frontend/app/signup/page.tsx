"use client";

import React, { useState } from "react";
import Link from "next/link";
import { useRouter } from "next/navigation";
import { useAuth } from "@/hooks/useAuth";
const SignupPage = () => {
  const {signup, isLoading} = useAuth()
  const router = useRouter();
 
  const [formData, setFormData] = useState({
    name: "",
    email: "",
    password: "",
    confirmPassword: "",
  });

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    if (formData.password !== formData.confirmPassword) {
      alert("Passwords do not match!");
      return;
    }
    signup(formData)
    console.log("Signing up:", formData);
  };

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setFormData({ ...formData, [e.target.name]: e.target.value });
  };

  return (
    <div className="min-h-screen bg-[#0f172a] flex flex-col items-center justify-center p-4">
      {/* Brand Logo */}
      <div className="flex items-center gap-2 mb-8">
        <div className="bg-white rounded-full p-1">
          <span className="text-xl">🐷</span>
        </div>
        <h1 className="text-2xl font-bold text-white tracking-tight">Piggy</h1>
      </div>

      {/* Signup Card */}
      <div className="w-full max-w-md bg-[#1e293b] border border-slate-700 rounded-xl shadow-2xl p-8">
        <h2 className="text-2xl font-semibold text-white mb-2 text-center">
          Create Account
        </h2>
        <p className="text-slate-400 text-sm text-center mb-8">
          Start tracking your savings and withdrawals today
        </p>

        <form onSubmit={handleSubmit} className="space-y-5">
          <div>
            <label className="block text-sm font-medium text-slate-300 mb-2">
              Full Name
            </label>
            <input
              name="name"
              type="text"
              required
              className="w-full bg-[#0f172a] border border-slate-600 rounded-lg py-3 px-4 text-white focus:outline-none focus:ring-2 focus:ring-blue-600 transition-all"
              placeholder="John Doe"
              onChange={handleChange}
              value={formData.name}
            />
          </div>

          <div>
            <label className="block text-sm font-medium text-slate-300 mb-2">
              Email Address
            </label>
            <input
              name="email"
              type="email"
              required
              className="w-full bg-[#0f172a] border border-slate-600 rounded-lg py-3 px-4 text-white focus:outline-none focus:ring-2 focus:ring-blue-600 transition-all"
              placeholder="name@example.com"
              onChange={handleChange}
              value={formData.email}
            />
          </div>

          <div>
            <label className="block text-sm font-medium text-slate-300 mb-2">
              Password
            </label>
            <input
              name="password"
              type="password"
              required
              className="w-full bg-[#0f172a] border border-slate-600 rounded-lg py-3 px-4 text-white focus:outline-none focus:ring-2 focus:ring-blue-600 transition-all"
              placeholder="••••••••"
              onChange={handleChange}
              value={formData.password}
            />
          </div>

          <div>
            <label className="block text-sm font-medium text-slate-300 mb-2">
              Confirm Password
            </label>
            <input
              name="confirmPassword"
              type="password"
              required
              className="w-full bg-[#0f172a] border border-slate-600 rounded-lg py-3 px-4 text-white focus:outline-none focus:ring-2 focus:ring-blue-600 transition-all"
              placeholder="••••••••"
              onChange={handleChange}
            />
          </div>

          <button
            type="submit"
            className="w-full bg-[#2563eb] hover:bg-[#1d4ed8] text-white font-semibold py-3 rounded-lg shadow-lg shadow-blue-900/20 transition-colors mt-2"
          >
            Create Account
          </button>
        </form>

        <div className="mt-8 text-center">
          <p className="text-slate-400 text-sm">
            Already have an account?{" "}
            <Link
              href="/login"
              className="text-blue-500 font-medium hover:underline"
            >
              Sign In
            </Link>
          </p>
        </div>
      </div>
    </div>
  );
};

export default SignupPage;
