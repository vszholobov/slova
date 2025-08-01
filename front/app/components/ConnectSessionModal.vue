<template>
  <div class="modal" @click.self="close">
    <div class="modal-content">
      <h3>Подключиться к сессии</h3>
      <input v-model="inputSessionId" placeholder="ID сессии" />
      <input v-model="username" placeholder="Имя пользователя" />
      <button @click="connectToGame" :disabled="loading">
        {{ loading ? 'Подключение...' : 'Подключиться' }}
      </button>
      <button @click="close">Отмена</button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useRuntimeConfig } from '#app';

const emit = defineEmits(['connected', 'close']);

const inputSessionId = ref('');
const username = ref('');
const loading = ref(false);
const config = useRuntimeConfig();

const connectToGame = () => {
  if (!inputSessionId.value.trim() || !username.value.trim()) {
    alert('Введите ID и имя!');
    return;
  }

  loading.value = true;

  const baseUrl = config.public.apiBaseUrl.replace(/\/$/, '');
  const wsUrl = `${config.public.apiWsSchema}${baseUrl}/session/${inputSessionId.value.trim()}`;

  const socket = new WebSocket(wsUrl);

  socket.onopen = () => {
    socket.send(
      JSON.stringify({
        command: 'SET_USERNAME',
        body: { username: username.value.trim() },
      })
    );
    emit('connected', socket);
    loading.value = false;
  };

  socket.onerror = (err) => {
    alert('Ошибка подключения к серверу');
    console.error(err);
    loading.value = false;
  };
};

const close = () => emit('close');
</script>

<style scoped>
.modal {
  position: fixed;
  top: 0; left: 0; width: 100%; height: 100%;
  background-color: rgba(0, 0, 0, 0.5);
  display: flex; justify-content: center; align-items: center;
}
.modal-content {
  background-color: white;
  padding: 20px;
  border-radius: 5px;
  text-align: center;
}
input {
  display: block;
  margin: 10px auto;
  padding: 8px;
  width: 80%;
  max-width: 300px;
}
button {
  margin: 5px;
  padding: 8px 16px;
}
</style>
