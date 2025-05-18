"use client";
import { useState, useEffect } from "react";

export default function Home() {
  const [data, setData] = useState("");
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const hostname = window.location.hostname;
    const backendUrl = `http://${hostname}:8080/db-version`;

    const fetchData = async () => {
      try {
        const response = await fetch(backendUrl);
        const data = await response.text();
        setData(data);
      } catch (err) {
        console.error("Error fetching data:", err);
        setError(err instanceof Error ? err.message : String(err));
      }
    };

    fetchData();
  }, []);

  return (
    <main className="font-mono">
      <div className="container-1600">
        <div className="py-4">
          <h1>Next.js with Go Backend</h1>
          <p>Response from Go server: {data}</p>
          {error && <p>Error: {error}</p>}
        </div>
      </div>
    </main>
  );
}
