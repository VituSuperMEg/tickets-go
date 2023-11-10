import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import { api } from "../../../services/api";
import { FaArrowRight, FaClock, FaTicketAlt } from "react-icons/fa";

interface ISession {
  id: string;
  image_path: string;
  session: string;
  description: string;
  film_count: number;
  film_time: number;
}

export function Session() {
  const [list, setList] = useState<ISession[]>([]);
  const { id } = useParams();
  useEffect(() => {
    async function getList() {
      const response = await api.get(`/films/${id}`);
      setList(response.data);
    }
    getList();
  }, [id]);
  return (
    <div className="session">
      <div className="session-card">
        <div className="content">
          <div className="session-content-card">
            <img src={list.image_path} alt="" />
            <div className="column">
              <h1 className="block">{list.session}</h1>
              <span className="mt-10" style={{
                width: 90
              }}>
                <span className="flex">
                  <FaClock color="#007BFF" size={20} />
                  &nbsp;&nbsp;
                  <span className="gray">
                   {list.film_time} min
                  </span>
                </span>
              </span>
              <span className="mt-10" style={{
                width: 200
              }}>
                <span className="flex">
                  <FaTicketAlt color="#007BFF" size={20} />
                  &nbsp;&nbsp;&nbsp;
                  <span className="gray">{list.film_count} sessões disponíveis</span>
                </span>
              </span>
              <div className="end-vertical">
              <button className="btn-btn-primary">
                Adicionar novas sessões
              </button>
              </div>
            </div>
          </div>
          <p className="gray">{list.description}</p>
          <div
            className="end"
            style={{
              marginTop: 90,
            }}
          >
            <FaArrowRight size={30} color="#858585" />
          </div>
        </div>
      </div>
      <div className="session-ticket"></div>
    </div>
  );
}
