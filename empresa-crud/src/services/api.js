import axios from 'axios';

const api = axios.create({
  baseURL: 'http://localhost:8080',
});

export const getCompanies = () => api.get('/empresa');
export const createCompany = (company) => api.post('/empresa', company);
export const updateCompany = (codigo, company) => api.put(`/empresa?codigo=${codigo}`, company);
export const deleteCompany = (codigo) => api.delete(`/empresa?codigo=${codigo}`);

export default api;
