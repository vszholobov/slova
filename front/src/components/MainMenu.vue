<template>
  <div class="max-w-md mx-auto">
    <div class="bg-white rounded-lg shadow-xl p-8">
      <h1 class="text-3xl font-bold text-center text-gray-800 mb-8">
        Скажи одно и то же
      </h1>
      
      <div class="space-y-4">
        <CreateSessionButton />
        <ConnectToSessionButton />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import CreateSessionButton from './CreateSessionButton.vue'
import ConnectToSessionButton from './ConnectToSessionButton.vue'
</script>