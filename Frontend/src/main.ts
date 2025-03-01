// main.js
import { createApp } from 'vue';
import App from './App.vue';
import router from './router';
import './index.css'; // Your main Tailwind CSS file
import './assets/fonts/fonts.css'; // Import your custom fonts

const app = createApp(App);

app.use(router);

app.mount('#app');