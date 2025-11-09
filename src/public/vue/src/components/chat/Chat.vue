<template>
  <div class="chat-container">
    <div class="messages" ref="messagesEnd">
      <div v-for="(msg, i) in messages" :key="i" class="message-container">
        <div class="message" :class="`message-${msg.user}`">{{ msg.text }}</div>
      </div>
    </div>
    <div class="chat-input">
      <input v-model="newMsg" @keyup.enter="sendMessage" placeholder="Pregunta tu weá…" />
      <button @click="sendMessage"><h2>↑</h2></button>
    </div>
  </div>
</template>
<script src="./Chat.js"></script>
<style src="./Chat.css"></style>
