<template>
  <div dir="rtl" class="min-h-screen flex items-center justify-center relative">
    <CommonPage>
      <AttendBox />
    </CommonPage>

    <teleport to="body">
      <div v-show="isModalVisible" class="modal-overlay">
        <div class="modal-content">
          <h2 class="text-lg font-semibold text-center mb-4">تسجيل الدخول</h2>
          <label for="modalUsername" class="block text-right mb-1">:اسم المستخدم</label>
          <input 
            type="text" 
            id="modalUsername" 
            v-model="modalUsername" 
            class="w-full p-2 border border-gray-300 rounded mb-4" 
            required
            @keydown.enter="focusNext('modalPassword')" 
          />
          
          <label for="modalPassword" class="block text-right mb-1">:كلمة المرور</label>
          <input 
            type="password" 
            id="modalPassword" 
            v-model="modalPassword" 
            class="w-full p-2 border border-gray-300 rounded mb-4" 
            required
            @keydown.enter="submitModal" 
          />
          
          <div class="flex justify-between">
            <button @click="submitModal" class="bg-[#5daaae] text-white px-4 py-2 rounded hover:bg-[#5daaae]">تسجيل الدخول</button>
          </div>
        </div>
      </div>
    </teleport>
  </div>
</template>

<script>
import { ref, computed } from 'vue';
import CommonPage from '../components/CommonPage.vue';
// import AttendBox from '../components/AttendE7twaa.vue';
import AttendBox from '../components/AttendOutE7twaa.vue';


export default {
  components: {
    CommonPage,
    AttendBox,
  },
  setup() {
    const modalUsername = ref('');
    const modalPassword = ref('');
    const isModalVisible = ref(false);
    const toastVisible = ref(false);
    const toastMessage = ref('');
    const loginFailed = ref(false); // New reactive property for login state

    const login = (username, password) => {
      const validUsername = 'keswah-attend';
      const validPassword = 'كسوة-فرح-نظام-آمن';

      if (username === validUsername && password === validPassword) {
        toastMessage.value = 'تسجيل الدخول ناجح! يتم توجيهك إلى الصفحة الرئيسية...';
        toastVisible.value = true;
        setTimeout(() => toastVisible.value = false, 3000);
        closeModal();
        loginFailed.value = false; // Reset on successful login
      } else {
        toastMessage.value = 'فشل تسجيل الدخول. يرجى التحقق من بيانات الاعتماد الخاصة بك.';
        toastVisible.value = true;
        setTimeout(() => toastVisible.value = false, 3000);
        loginFailed.value = true; // Set to true on login failure
      }
    };

    const openModal = () => {
       if (!isModalAlreadyOpened()) {
        isModalVisible.value = true;

       }
    };

    const closeModal = () => {
      isModalVisible.value = false;
    };

    const submitModal = () => {
      login(modalUsername.value, modalPassword.value);
    };

    const isModalAlreadyOpened = () => {
      return localStorage.getItem('modalOpened') === 'true';
    };

    const setModalOpenedFlag = () => {
      localStorage.setItem('modalOpened', 'true');
    };

    // Computed property to determine the toast color
    const toastClass = computed(() => {
      return loginFailed.value ? 'bg-red-500' : 'bg-green-500'; // Red for failure, green for success
    });

    // Method to focus the next input field
    const focusNext = (nextFieldId) => {
      const nextField = document.getElementById(nextFieldId);
      if (nextField) {
        nextField.focus();
      }
    };

    return {
      modalUsername,
      modalPassword,
      isModalVisible,
      toastVisible,
      toastMessage,
      openModal,
      closeModal,
      submitModal,
      setModalOpenedFlag,
      loginFailed,
      toastClass, // Include computed property
      focusNext, // Include focusNext method
    };
  },
  mounted() {
    this.openModal();
  },
};
</script>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Mirza:wght@400;500;600;700&display=swap');
.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: rgba(0, 0, 0, 1); 
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 1000; 
}

.modal-content {
  background-color: #ffffff; 
  padding: 20px;
  width: 90%;
  max-width: 400px; 
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1); 
}



.fixed.inset-0 {
  background-color: rgba(0, 0, 0, 0.7); /* خلفية مظلمة مع شفافية */
  z-index: 50; /* تأكد من أن المودال في المقدمة */
}

.bg-opacity-70 {
  background-color: rgba(0, 0, 0, 0.7);
}
.fade-enter-active, .fade-leave-active {
  transition: opacity 0.5s;
}
.fade-enter, .fade-leave-to {
  opacity: 0;
}

</style>