import { FaInfoCircle, FaClock } from "react-icons/fa";
import { useState, useEffect } from "react";
import { api } from "../../services/api";
import { useNavigate } from "react-router-dom";

interface ICardFilme {
  id : string;
  image_path : string;
  session : string;
  description : string;
  film_count : number;
  film_time : number;
}
export function CardFilme() {
  const [list, setList] = useState<ICardFilme[]>([]);
  const navigate = useNavigate()

  useEffect(() => {
    async function getList() {
      const response = await api.get("/films");
      setList(response.data);
    }
    getList();
  }, []);
  
  function handleLotacao(id : string) {
    navigate(`/lotacao/${id}`);
  }
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
            <button className="btn-btn-info-small" 
            onClick={() => handleLotacao(i.id)}>
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
