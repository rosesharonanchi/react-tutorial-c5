"use client";

import React, { useState } from "react";
import Link from "next/link";
import { useAuth } from "@/hooks/useAuth";
const LoginPage = () => {

  const { login, isLoading } = useAuth();

  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    // Handle authentication logic here
    await login({ email, password });
    console.log("Logging in with:", { email, password });
  };

  return (
    <div className="min-h-screen bg-[#0f172a] flex flex-col items-center justify-center p-4">
      {/* Logo Section */}
      <div className="flex items-center gap-2 mb-8">
        <div className="bg-white rounded-full p-1">
          <span className="text-xl">🐷</span>
        </div>
        <h1 className="text-2xl font-bold text-white tracking-tight">Piggy</h1>
      </div>

      {/* Login Card */}
      <div className="w-full max-w-md bg-[#1e293b] border border-slate-700 rounded-xl shadow-2xl p-8">
        <h2 className="text-2xl font-semibold text-white mb-2 text-center">
          Welcome Back
        </h2>
        <p className="text-slate-400 text-sm text-center mb-8">
          Enter your details to manage your savings
        </p>

        <form onSubmit={handleSubmit} className="space-y-6">
          <div>
            <label className="block text-sm font-medium text-slate-300 mb-2">
              Email Address
            </label>
            <input
              type="email"
              required
              className="w-full bg-[#0f172a] border border-slate-600 rounded-lg py-3 px-4 text-white placeholder-slate-500 focus:outline-none focus:ring-2 focus:ring-blue-600 focus:border-transparent transition-all"
              placeholder="name@example.com"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
            />
          </div>

          <div>
            <div className="flex justify-between mb-2">
              <label className="text-sm font-medium text-slate-300">
                Password
              </label>
              <a href="#" className="text-xs text-blue-500 hover:underline">
                Forgot password?
              </a>
            </div>
            <input
              type="password"
              required
              className="w-full bg-[#0f172a] border border-slate-600 rounded-lg py-3 px-4 text-white placeholder-slate-500 focus:outline-none focus:ring-2 focus:ring-blue-600 focus:border-transparent transition-all"
              placeholder="••••••••"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
            />
          </div>

          <button
            type="submit"
            className="w-full bg-[#2563eb] hover:bg-[#1d4ed8] text-white font-semibold py-3 rounded-lg shadow-lg shadow-blue-900/20 transition-colors"
          >
            Sign In
          </button>
        </form>

        <div className="mt-8 text-center">
          <p className="text-slate-400 text-sm">
            Don't have an account?{" "}
            <Link
              href="/signup"
              className="text-blue-500 font-medium hover:underline"
            >
              Create one
            </Link>
          </p>
        </div>
      </div>

      {/* Footer info */}
      <p className="mt-8 text-slate-500 text-xs">
        &copy; {new Date().getFullYear()} Piggy Finance. All rights reserved.
      </p>
    </div>
  );
};

export default LoginPage;
