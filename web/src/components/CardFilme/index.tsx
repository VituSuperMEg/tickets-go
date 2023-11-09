import { FaInfoCircle, FaClock } from "react-icons/fa";
import { useState, useEffect } from "react";
import { api } from "../../services/api";

export function CardFilme() {
  const [list, setList] = useState([]);

  useEffect(() => {
    async function getList() {
      const response = await api.get("/films");
      setList(response.data);
    }
    getList();
  }, []);

  return (
    <div className="dashboard">
      {list.map((i) => (
        <div key={i.id} className="card-filme">
          <img src={`${i.image_path}`} alt="" width={200} />
          <h1 className="block">{i.session}</h1>
          <p className="description">{i.description}</p>
          <div className="flex">
            <p className="block">
              Sessões : <span className="maximum">{i.film_count}</span>
            </p>
            <button className="btn-btn-info-small">
              <FaInfoCircle size={20} /> Visualizar Lotação
            </button>
          </div>
          <div className="end" style={{
            marginTop : -10
          }}>
            <FaClock color="#007BFF" />
            <span className="gray">{i.film_time}</span>
          </div>
        </div>
      ))}
    </div>
  );
}
