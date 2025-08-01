<template>
  <div class="container">
    <h1>Добро пожаловать в игру в слова!</h1>
    <button @click="createSession">Создать сессию</button>
    <div v-if="sessionId">
      <p>Сессия создана! ID: {{ sessionId }}</p>
    </div>
    <button @click="connectToSession">Подключиться к сессии</button>
  </div>
</template>

<script lang="ts">
import { ref } from 'vue';
import { useRuntimeConfig } from '#app';

interface CreateSessionResponse {
  id: string;
}

export default {
  setup() {
    const sessionId = ref<string | null>(null);
    const config = useRuntimeConfig();

    const createSession = async () => {
      try {
        const { data } = await useFetch<CreateSessionResponse>(`${config.public.apiBaseUrl}api/create-session`, {
          method: 'POST',
        });
        sessionId.value = data.value?.id || null;
      } catch (error) {
        console.error('Ошибка при создании сессии:', error);
      }
    };

    const connectToSession = () => {
      console.log('Подключение к существующей сессии');
    };

    return {
      sessionId,
      createSession,
      connectToSession,
    };
  },
};
</script>

<style scoped>
.container {
  max-width: 600px;
  margin: 0 auto;
  text-align: center;
}
button {
  margin: 10px;
}
</style>