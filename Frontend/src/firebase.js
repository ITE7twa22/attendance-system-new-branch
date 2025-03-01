// firebase.js

import { initializeApp } from 'firebase/app';
import { getAuth } from 'firebase/auth';

const firebaseConfig = {
  apiKey: 'AIzaSyBHEk5P7FMuM8CSYZ5DI75l0PZZOEkHdMo',
  authDomain: "volunteersdata-cf17b.firebaseapp.com",
  projectId: "volunteersdata-cf17b",
  storageBucket: "volunteersdata-cf17b.appspot.com",
  messagingSenderId: "542427088882",
  appId: "1:542427088882:web:28637e5637ba5f20fc8d6b",
};

const firebaseApp = initializeApp(firebaseConfig);
const auth = getAuth(firebaseApp);

export default auth;
