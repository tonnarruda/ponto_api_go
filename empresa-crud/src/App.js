import React, { useState } from 'react';
import './index.css'; // Importe seu arquivo de estilos CSS aqui
import CompanyEdit from './components/companyEdit';
import CompanyList from './components/companyList';
import CompanyForm from './components/companyForm';

const App = () => {
  const [selectedCompany, setSelectedCompany] = useState(null);

  const handleEdit = (company) => {
    setSelectedCompany(company);
  };

  const handleUpdate = () => {
    setSelectedCompany(null);
  };

  return (
    <div className="app-container">
      <h1>Company Management</h1>
      <div className="company-form">
        <CompanyForm onAdd={handleUpdate} />
      </div>
      <ul className="company-list">
        <CompanyList onEdit={handleEdit} />
      </ul>
      {selectedCompany && <CompanyEdit company={selectedCompany} onUpdate={handleUpdate} />}
    </div>
  );
};

export default App;
