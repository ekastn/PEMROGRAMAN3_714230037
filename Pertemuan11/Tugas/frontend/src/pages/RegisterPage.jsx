import { useState } from "react";
import { register } from "../services/authService";
import Swal from "sweetalert2";
import { useNavigate, Link } from "react-router-dom";
import { useEffect } from "react";

export function RegisterPage() {
    const [form, setForm] = useState({ username: "", password: "", role: "" });
    const navigate = useNavigate();

    useEffect(() => {
        const token = localStorage.getItem("token");
        if (token) navigate("/dashboard");
    }, []);

    const handleChange = (e) => {
        const { name, value } = e.target;
        setForm((prev) => ({ ...prev, [name]: value }));
    };

    const handleSubmit = async (e) => {
        e.preventDefault();

        if (!form.username || !form.password || !form.role) {
            Swal.fire("Gagal", "Semua field harus diisi", "error");
            return;
        }

        if (form.role !== "admin" && form.role !== "user") {
            Swal.fire("Gagal", "Role harus admin atau user", "error");
            return;
        }

        try {
            await register(form.username, form.password, form.role);
            Swal.fire({
                icon: "success",
                title: "Berhasil!",
                text: "Berhasil membuat user",
                confirmButtonColor: "#22c55e", // Tailwind green-500
            }).then(() => {
                navigate("/");
            });
        } catch (err) {
            Swal.fire("Gagal", "Gagal membuat user", "error");
        }
    };

    return (
        <div className="min-h-screen flex justify-center items-center bg-gray-100">
            <form
                onSubmit={handleSubmit}
                className="bg-white p-8 rounded shadow-md w-full max-w-sm"
            >
                <h2 className="text-xl font-bold mb-6 text-center">Login</h2>
                <input
                    type="text"
                    name="username"
                    placeholder="Username"
                    value={form.username}
                    onChange={handleChange}
                    className="w-full p-2 mb-4 border rounded"
                />
                <input
                    type="password"
                    name="password"
                    placeholder="Password"
                    value={form.password}
                    onChange={handleChange}
                    className="w-full p-2 mb-6 border rounded"
                />
                <select
                    name="role"
                    value={form.role}
                    onChange={handleChange}
                    className="w-full p-2 mb-6 border rounded"
                >
                    <option value="" disabled>Pilih Role</option>
                    <option value="admin">Admin</option>
                    <option value="user">User</option>
                </select>
                <button
                    type="submit"
                    className="w-full bg-blue-500 text-white py-2 rounded hover:bg-blue-600"
                >
                    Login
                </button>
                <p className="mt-4 text-center">
                    Sudah punya akun?{" "}
                    <Link to="/" className="text-blue-500">
                        Login
                    </Link>
                </p>
            </form>
        </div>
    );
}
