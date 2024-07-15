import React, { useState, useEffect } from 'react';
import { getCompanies, deleteCompany } from '../services/api';

const CompanyList = ({ onEdit }) => {
  const [companies, setCompanies] = useState([]);

  useEffect(() => {
    fetchCompanies();
  }, []);

  const fetchCompanies = async () => {
    const response = await getCompanies();
    setCompanies(response.data);
  };

  const handleDelete = async (codigo) => {
    await deleteCompany(codigo);
    fetchCompanies();
  };

  return (
    <div>
      <h2>Companies List</h2>
      <ul>
        {companies.map((company) => (
          <li key={company.codigo}>
            {company.nome} - {company.razao_social}
            <button onClick={() => onEdit(company)}>Edit</button>
            <button onClick={() => handleDelete(company.codigo)}>Delete</button>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default CompanyList;
