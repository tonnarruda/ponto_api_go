import React, { useState } from 'react';
import { createCompany } from '../services/api';

const CompanyForm = ({ onAdd }) => {
  const [company, setCompany] = useState({
    nome: '',
    razao_social: '',
    cnpj_base: '',
    convert_tipo_he: 1,
    cpf: '',
    dt_encerramento: '',
    ultima_atualizacao_ac: '',
    falta_ajustar_no_ac: 0,
    aderiu_esocial: 1,
    data_adesao_esocial: '',
    data_adesao_esocial_f2: '',
    tp_amb_esocial: 2,
    status_envio_app: 1,
    nmfantasia: '',
    cnpj_licenciado: '',
    freemium_last_update: '0',
  });

  const handleChange = (e) => {
    const { name, value } = e.target;
    setCompany({ ...company, [name]: value });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      await createCompany(company);
      onAdd();
      setCompany({
        nome: '',
        razao_social: '',
        cnpj_base: '',
        convert_tipo_he: 1,
        cpf: '',
        dt_encerramento: '',
        ultima_atualizacao_ac: '',
        falta_ajustar_no_ac: 0,
        aderiu_esocial: 1,
        data_adesao_esocial: '',
        data_adesao_esocial_f2: '',
        tp_amb_esocial: 2,
        status_envio_app: 1,
        nmfantasia: '',
        cnpj_licenciado: '',
        freemium_last_update: '0',
      });
    } catch (err) {
      console.error(err);
    }
  };

  return (
    <form onSubmit={handleSubmit}>
      <h2>Create Company</h2>
      <input
        name="nome"
        placeholder="Nome"
        value={company.nome}
        onChange={handleChange}
      />
      <input
        name="razao_social"
        placeholder="Razao Social"
        value={company.razao_social}
        onChange={handleChange}
      />
      <input
        name="cnpj_base"
        placeholder="CNPJ Base"
        value={company.cnpj_base}
        onChange={handleChange}
      />
      <input
        name="cpf"
        placeholder="CPF"
        value={company.cpf}
        onChange={handleChange}
      />
      <input
        name="dt_encerramento"
        placeholder="Data Encerramento"
        value={company.dt_encerramento}
        onChange={handleChange}
      />
      <input
        name="ultima_atualizacao_ac"
        placeholder="Ultima Atualizacao AC"
        value={company.ultima_atualizacao_ac}
        onChange={handleChange}
      />
      <input
        name="data_adesao_esocial"
        placeholder="Data Adesao eSocial"
        value={company.data_adesao_esocial}
        onChange={handleChange}
      />
      <input
        name="data_adesao_esocial_f2"
        placeholder="Data Adesao eSocial F2"
        value={company.data_adesao_esocial_f2}
        onChange={handleChange}
      />
      <input
        name="nmfantasia"
        placeholder="Nome Fantasia"
        value={company.nmfantasia}
        onChange={handleChange}
      />
      <input
        name="cnpj_licenciado"
        placeholder="CNPJ Licenciado"
        value={company.cnpj_licenciado}
        onChange={handleChange}
      />
      <button type="submit">Create</button>
    </form>
  );
};

export default CompanyForm;
