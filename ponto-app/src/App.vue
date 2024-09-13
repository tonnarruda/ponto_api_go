<template>
  <div class="login-container">
    <h2>Login</h2>
    <form @submit.prevent="login">
      <div class="form-group">
        <label for="username">Usuário:</label>
        <input type="text" v-model="username" id="username" required />
      </div>
      <div class="form-group">
        <label for="password">Senha:</label>
        <input type="password" v-model="password" id="password" required />
      </div>
      <div class="form-group">
        <label for="company">Empresa:</label>
        <select v-model="selectedCompany" id="company" required>
          <option v-for="company in companies" :key="company.id" :value="company.id">
            {{ company.nome }}
          </option>
        </select>
      </div>
      <button type="submit">Entrar</button>
    </form>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      username: '',
      password: '',
      selectedCompany: '',
      companies: []
    };
  },
  methods: {
    async fetchCompanies() {
      try {
        const response = await axios.get('http://localhost:8080/empresa');
        this.companies = response.data;
      } catch (error) {
        console.error("Erro ao buscar empresas:", error);
      }
    },
    login() {
      console.log('Usuário:', this.username);
      console.log('Senha:', this.password);
      console.log('Empresa:', this.selectedCompany);
      // Aqui você pode adicionar a lógica de login real
    }
  },
  created() {
    this.fetchCompanies(); // Busca a lista de empresas ao carregar o componente
  }
};
</script>

<style scoped>
.login-container {
  max-width: 400px;
  margin: 100px auto;
  padding: 20px;
  background-color: #f5f5f5;
  border-radius: 8px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
}

h2 {
  text-align: center;
  margin-bottom: 20px;
}

.form-group {
  margin-bottom: 15px;
}

label {
  display: block;
  margin-bottom: 5px;
}

input, select {
  width: 100%;
  padding: 10px;
  margin-bottom: 10px;
  border: 1px solid #ccc;
  border-radius: 4px;
}

button {
  width: 100%;
  padding: 10px;
  background-color: #4CAF50;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

button:hover {
  background-color: #45a049;
}
</style>
