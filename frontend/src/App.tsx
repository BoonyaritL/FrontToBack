import { useEffect, useState } from 'react';

// Type ให้ตรงกับ Rust
interface Todo {
  id: number;
  title: string;
  completed: boolean;
}

function App() {
  const [todos, setTodos] = useState<Todo[]>([]);
  const [input, setInput] = useState("");
  const API_URL = "http://localhost:3000/todos";

  // 1. Read (ดึงข้อมูลเมื่อเปิดเว็บ)
  useEffect(() => {
    fetchTodos();
  }, []);

  const fetchTodos = async () => {
    const res = await fetch(API_URL);
    const data = await res.json();
    setTodos(data);
  };

  // 2. Create (เพิ่มข้อมูล)
  const addTodo = async () => {
    if (!input) return;
    await fetch(API_URL, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ title: input }),
    });
    setInput("");
    fetchTodos(); // ดึงใหม่เพื่อให้หน้าจออัปเดต
  };

  // 3. Update (ติ๊กถูก)
  const toggleTodo = async (id: number, currentStatus: boolean) => {
    await fetch(`${API_URL}/${id}`, {
      method: "PATCH",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ completed: !currentStatus }),
    });
    fetchTodos();
  };

  // 4. Delete (ลบ)
  const deleteTodo = async (id: number) => {
    await fetch(`${API_URL}/${id}`, {
      method: "DELETE",
    });
    fetchTodos();
  };

  return (
    <div style={{ padding: "2rem", fontFamily: "sans-serif", maxWidth: "500px", margin: "0 auto" }}>
      <h1>🦀 Rust + React CRUD ⚛️</h1>

      <div style={{ display: "flex", gap: "10px", marginBottom: "20px" }}>
        <input 
          value={input}
          onChange={(e) => setInput(e.target.value)}
          placeholder="ทำอะไรดีวันนี้..."
          style={{ flex: 1, padding: "8px" }}
          onKeyDown={(e) => e.key === "Enter" && addTodo()}
        />
        <button onClick={addTodo} style={{ padding: "8px 16px" }}>Add</button>
      </div>

      <ul style={{ listStyle: "none", padding: 0 }}>
        {todos.map((todo) => (
          <li key={todo.id} style={{ 
            display: "flex", 
            justifyContent: "space-between", 
            alignItems: "center",
            padding: "10px",
            borderBottom: "1px solid #eee",
            background: todo.completed ? "#f9f9f9" : "white"
          }}>
            <span 
              onClick={() => toggleTodo(todo.id, todo.completed)}
              style={{ 
                cursor: "pointer", 
                textDecoration: todo.completed ? "line-through" : "none",
                color: todo.completed ? "gray" : "black"
              }}
            >
              {todo.completed ? "✅" : "⬜"} {todo.title}
            </span>
            <button 
              onClick={() => deleteTodo(todo.id)}
              style={{ background: "red", color: "white", border: "none", padding: "5px 10px", cursor: "pointer", borderRadius: "4px" }}
            >
              Delete
            </button>
          </li>
        ))}
      </ul>
    </div>
  );
}

export default App;