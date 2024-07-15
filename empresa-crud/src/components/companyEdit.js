import React, { useState, useEffect } from 'react';
import { updateCompany } from '../services/api';

const CompanyEdit = ({ company, onUpdate }) => {
  const [updatedCompany, setUpdatedCompany] = useState(company);

  useEffect(() => {
    setUpdatedCompany(company);
  }, [company]);

  const handleChange = (e) => {
    const { name, value } = e.target;
    setUpdatedCompany({ ...updatedCompany, [name]: value });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      await updateCompany(updatedCompany.codigo, updatedCompany);
      onUpdate();
    } catch (err) {
      console.error(err);
    }
  };

  if (!company) return null;

  return (
    <form onSubmit={handleSubmit}>
      <h2>Edit Company</h2>
      <input
        name="nome"
        placeholder="Nome"
        value={updatedCompany.nome}
        onChange={handleChange}
      />
      <input
        name="razao_social"
        placeholder="Razao Social"
        value={updatedCompany.razao_social}
        onChange={handleChange}
      />
      <input
        name="cnpj_base"
        placeholder="CNPJ Base"
        value={updatedCompany.cnpj_base}
        onChange={handleChange}
      />
      <input
        name="cpf"
        placeholder="CPF"
        value={updatedCompany.cpf}
        onChange={handleChange}
      />
      <input
        name="dt_encerramento"
        placeholder="Data Encerramento"
        value={updatedCompany.dt_encerramento}
        onChange={handleChange}
      />
      <input
        name="ultima_atualizacao_ac"
        placeholder="Ultima Atualizacao AC"
        value={updatedCompany.ultima_atualizacao_ac}
        onChange={handleChange}
      />
      <input
        name="data_adesao_esocial"
        placeholder="Data Adesao eSocial"
        value={updatedCompany.data_adesao_esocial}
        onChange={handleChange}
      />
      <input
        name="data_adesao_esocial_f2"
        placeholder="Data Adesao eSocial F2"
        value={updatedCompany.data_adesao_esocial_f2}
        onChange={handleChange}
      />
      <input
        name="nmfantasia"
        placeholder="Nome Fantasia"
        value={updatedCompany.nmfantasia}
        onChange={handleChange}
      />
      <input
        name="cnpj_licenciado"
        placeholder="CNPJ Licenciado"
        value={updatedCompany.cnpj_licenciado}
        onChange={handleChange}
      />
      <button type="submit">Update</button>
    </form>
  );
};

export default CompanyEdit;
