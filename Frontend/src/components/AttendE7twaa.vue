<template>
  <div class="flex flex-col items-center justify-center">
    <div class="loader-container" id="loadingScreen">
      <div class="gear">
        <img src="../assets/img/Loader_Image.png" alt="" />
      </div>
    </div>

    <div v-if="!isConfirmationVisible && !isLogoutConfirmationVisible">
  <div class="focus-within:shadow-e7twa22-cyan/70 group flex flex-col md:flex-row items-center p-4 transition duration-300 overflow-clip shadow-lg" dir="auto">
    
    <input
      @keyup.enter.prevent="getRelationsInfo"
      @input="handleInputChange"
      v-model="volunteerID"
      type="text"
      placeholder="الرجاء إدخال رقم التسجيل أو الهوية هنا"
      class="border border-gray-300 border-dashed px-8 py-2 rounded-lg w-full shadow-md shadow-[#5daaae49] focus:outline-none focus:ring-2 focus:ring-[#5daaae] text-center ml-4"
      inputmode="numeric"
    />

    <div class="flex space-x-2 mt-4 md:mt-0 md:ml-2"> <!-- Space between buttons -->
      <button @click="getRelationsInfo" v-if="isSearch" class="btn text-white px-4 py-2 rounded-md cursor-pointer relative " style="border-radius: 12px;">التالي  </button>
      <button @click="handleCheckIn" v-if="enableButton" class="btn text-white px-4 py-2 rounded-md cursor-pointer relative"  style="border-radius: 12px;">{{ checkInButtonText }}</button>
    </div>
    


    
  </div>
</div>

    <div v-if="isConfirmationVisible" class="text-center">
      <p class="text-4xl text-white font-light" style="font-family: 'Mizan AR+LT';">أهلًا بك <span class="text-e7twa22-blue">{{ relationsVolunteerName }}</span> هل تريد تأكيد تسجيل الدخول؟</p>
   <div class="button-container">
    <button @click="confirmLogin" class="bg-green-500 text-white px-4 py-2 rounded-md cursor-pointer">نعم</button>
    <button @click="cancelLogin" class="bg-gray-500 text-white px-4 py-2 rounded-md cursor-pointer">إلغاء</button>
</div>
    </div>

    <div v-if="isLogoutConfirmationVisible" class="text-center">
      <p class="text-4xl text-white" style="font-family: 'Mizan AR+LT';">أهلًا بك <span class="text-e7twa22-blue ">{{ relationsVolunteerName }}</span> هل تريد تأكيد تسجيل الخروج؟</p>
      <div class="button-container">
    <button @click="confirmLogout" class="bg-red-500 text-white px-4 py-2 rounded-md cursor-pointer">نعم</button>
    <button @click="cancelLogout" class="bg-gray-500 text-white px-4 py-2 rounded-md cursor-pointer">إلغاء</button>
</div>
    </div>

    <br><br>

    <p v-if="isFindName" class="text-4xl text-white font-light" style="font-family: 'Mizan AR+LT';"> أهلا بك <span class="text-e7twa22-blue font-normal">{{ relationsVolunteerName }}</span> الرجاء تسجيل الدخول ✨!</p>
    <p v-if="ifSucsLogin" class="text-4xl text-white font-light" style="font-family: 'Mizan AR+LT';"> شكر الله سعيك <span class="text-e7twa22-blue font-normal">{{ relationsVolunteerName }}</span> ✨!</p>
    <p v-if="isLogin && !ifSucsLogin" class="text-4xl text-gray-500 font-light" style="font-family: 'Mizan AR+LT';"> {{ checkInOutPText }} <span class="text-e7twa22-blue font-normal">{{ relationsVolunteerName }}✨!</span></p>

    <div class="text-center flex flex-col gap-4">
      <p v-if="catchName" class="text-4xl text-white font-light" style="font-family: 'Mizan AR+LT';">{{ catchName }}</p>
    </div>

    <!-- Modal for Success/Error Messages -->
<!-- Modal for Success/Error Messages -->
<div 
  id="messageModal" 
  tabindex="-1" 
  aria-hidden="true" 
  :class="{'hidden': !modalVisible, 'flex': modalVisible}" 
  class="fixed inset-0 z-50 flex justify-center items-center"
>
  <div class="relative p-4 w-full max-w-xs mx-auto">
    <div 
      class="relative p-4 text-center rounded-lg shadow-lg"
      :class="isSuccess ? 'bg-white' : 'bg-white'"
    >
      <button 
        id="closeModal" 
        @click="closeModal" 
        class="text-gray-400 absolute top-2 right-2 bg-transparent hover:bg-gray-200 hover:text-gray-900 rounded-lg text-sm p-1.5"
      >
        <svg 
          aria-hidden="true" 
          class="w-5 h-5" 
          fill="currentColor" 
          viewBox="0 0 20 20" 
          xmlns="http://www.w3.org/2000/svg"
        >
          <path 
            fill-rule="evenodd" 
            d="M4.293 4.293a1 1 0 011.414 0L10 8.586l4.293-4.293a1 1 0 111.414 1.414L11.414 10l4.293 4.293a1 1 0 01-1.414 1.414L10 11.414l-4.293 4.293a1 1 0 01-1.414-1.414L8.586 10 4.293 5.707a1 1 0 010-1.414z" 
            clip-rule="evenodd"
          ></path>
        </svg>
        <span class="sr-only">Close modal</span>
      </button>
      <div 
        id="modalIcon" 
        class="w-12 h-12 rounded-full p-1 flex items-center justify-center mx-auto mb-3.5" 
        :class="isSuccess ? 'bg-green-200' : 'bg-red-200'"
      > 
        <svg 
          id="modalSvg" 
          aria-hidden="true" 
          class="w-8 h-8" 
          fill="currentColor" 
          viewBox="0 0 20 20" 
          xmlns="http://www.w3.org/2000/svg"
        >
          <path 
            id="successPath" 
            fill-rule="evenodd" 
            d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z" 
            clip-rule="evenodd" 
            v-if="isSuccess" 
            class="text-green-500"
          ></path>
          <path 
            id="errorPath" 
            fill-rule="evenodd" 
            d="M12 2C6.48 2 2 6.48 2 12c0 5.52 4.48 10 10 10s10-4.48 10-10S17.52 2 12 2zm1 15h-2v-2h2v2zm0-4h-2V7h2v6z" 
            clip-rule="evenodd" 
            v-if="!isSuccess" 
            class="text-red-500"
          ></path>
        </svg>
        <span class="sr-only"></span>
      </div>
      <p 
        id="modalMessage" 
        class="mb-2 text-2xl font-semibold text-center text-gray-900"
      >
        {{ modalMessage }}
      </p>
    </div>
  </div>
</div>
</div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      enableButton: false,
      isSearch: true,
      volunteerID: '',
      relationsVolunteerName: '',
      catchName: '',
      isLogin: false,
      isFindName: false,
      ifCheckIn: false,
      ifSucsLogin: false,
      isConfirmationVisible: false,
      isLogoutConfirmationVisible: false, // New state for logout confirmation
      modalMessage: '',
      modalVisible: false,
      isSuccess: false,
    };
  },
  computed: {
    checkInButtonText() {
      return this.ifCheckIn ? 'خروج' : 'دخول';
    },
    checkInOutPText() {
      return this.ifSucsLogin ? 'تم تسجيل دخول ' : 'رافقتك السلامة';
    },
  },
  methods: {
    handleInputChange() {
      // Reset states for new input
      this.catchName = "";
      this.enableButton = false;
      this.isSearch = true;
      this.ifSucsLogin = false;
      this.isFindName = false;
      this.relationsVolunteerName = "";
      this.ifCheckIn = false;
    },

    getRelationsInfo() {
      this.showLoadingScreen();
      axios.post('/getRelationsInfo', { volunteerID: parseInt(this.volunteerID) }, {
        headers: {
          'Content-Type': 'application/json',
        },
      })
      .then(response => {
        this.relationsVolunteerName = response.data.RelationsVolunteerName;
        this.isSearch = false;
        this.isConfirmationVisible = true; // Show confirmation message for login
        this.getIfCheckIn(); // Check if the volunteer is already checked in
      })
      .catch(error => {
        this.relationsVolunteerName = "";
        this.isFindName = false;
        this.catchName = "الكود المدخل غير صحيح";
      })
      .finally(() => {
        this.hideLoadingScreen();
      });
    },

    confirmLogin() {
      this.isConfirmationVisible = false; // Hide confirmation message
      this.handleCheckIn(); // Call existing method to handle check-in
    },

    cancelLogin() {
      this.isConfirmationVisible = false; // Hide confirmation message
      this.resetStates(); // Reset states to allow fresh input
    },

    

    cancelLogout() {
      this.isLogoutConfirmationVisible = false; // Hide logout confirmation message
    },

    resetStates() {
      this.volunteerID = ''; 
      this.relationsVolunteerName = ''; 
      this.isSearch = true; 
      this.ifSucsLogin = false; 
      this.isFindName = false; 
      this.ifCheckIn = false;
      this.isLogin = false; // Reset login state
      this.catchName = ''; // Clear catch name
    },

    handleCheckIn() {
    this.showLoadingScreen();
    
    // Determine endpoint based on check-in state
    const endpoint = this.ifCheckIn ? '/updateLogoutDateTime' : '/addRelationsVolunteer';
    const loginDateTime = new Date();

    // If already checked in, request logout confirmation
    if (this.ifCheckIn) {
      this.requestLogout(); // Show logout confirmation instead of handling check-out directly
      this.hideLoadingScreen();
      return; // Early exit
    }

    // Proceed to check-in logic
    axios.post(endpoint, {
      volunteerID: parseInt(this.volunteerID),
      RelationsVolunteerName: this.relationsVolunteerName,
      LoginDateTime: loginDateTime.toISOString(),
    }, {
      headers: {
        'Content-Type': 'application/json',
      },
    })
    .then(response => {
      this.enableButton = false;
      this.isSearch = true;

      // Check if logging in or out
      if (this.ifCheckIn) {
        this.showMessageModal(true, "تم تسجيل الخروج بنجاح");
      } else {
        this.showMessageModal(true, "تم تسجيل الدخول بنجاح");
      }

      this.resetStates(); // Reset states after login/logout
    })
    .catch(error => {
      console.log("ERROR!!!!", error);
      this.showMessageModal(false, "حدث خطأ، يرجى المحاولة مرة أخرى.");
    })
    .finally(() => {
      this.hideLoadingScreen();
    });
  },

  requestLogout() {
    this.isLogoutConfirmationVisible = true; // Show logout confirmation message
  },

  confirmLogout() {
    this.isLogoutConfirmationVisible = false; // Hide logout confirmation
    this.ifCheckIn = false; // Mark as logged out
    this.showMessageModal(true, "تم تسجيل الخروج بنجاح"); // Show success message
    this.resetStates(); // Reset states after logout
  },

    async getIfCheckIn() {
      try {
        const response = await axios.post(
          '/searchAttendanceHandler',
          { volunteerID: parseInt(this.volunteerID) },
          {
            headers: {
              'Content-Type': 'application/json',
            },
          }
        );

        if (response.data === 'Volunteer not came today') {
          this.ifCheckIn = false;
        } else if (response.data === 'Volunteer came today') {
          this.ifCheckIn = true;
        } else {
          console.error('Unexpected response:', response.data);
        }
      } catch (error) {
        console.error('Error:', error.message || 'An error occurred');
      }
    },

    showLoadingScreen() {
      const loadingScreen = document.getElementById('loadingScreen');
      if (loadingScreen) {
        loadingScreen.style.display = 'flex'; // Show loading screen
      }
    },

    hideLoadingScreen() {
      const loadingScreen = document.getElementById('loadingScreen');
      if (loadingScreen) {
        loadingScreen.style.display = 'none'; // Hide loading screen
      }
    },

    showMessageModal(isSuccess, message) {
      this.modalMessage = message;
      this.modalVisible = true; // Show modal
      this.isSuccess = isSuccess; // Set success state

      // Automatically close the modal after 5 seconds
      setTimeout(() => {
        this.closeModal();
      }, 1500);
    },

    closeModal() {
      this.modalVisible = false; // Hide modal
      this.isSuccess = false; // Reset success state
    }
  },
};
</script>

<style>
.fixed-width-input {
  width: 300px;
}

.text-e7twa22-blue {

  --tw-text-opacity: 1;

  color: rgb(93 170 174 / var(--tw-text-opacity));
  font-weight: 5000;

}

#loadingScreen {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(93, 170, 174, 0.8);
  color: white;
  display: flex;
  justify-content: center;
  align-items: center;
  font-size: 2em;
  z-index: 9999;
  display: none; /* Hidden by default */
}

.loader-container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
}
.button-container {
  display: flex;
  justify-content: center; /* Center the buttons */
  gap: 16px; /* Adjust space between buttons */
  padding: 10px;
}
.btn{
  background-color: #5daaae;
}
.gear {
  width: 100px;
  height: 100px;
  animation: rotate 4s linear infinite;
}

.gear img {
  width: 100%;
  height: 100%;
  filter: brightness(96%) contrast(98%);
}


@keyframes rotate {
  from {
    transform: rotate(0deg);
  }
  to {
    transform: rotate(360deg);
  }
}
*{
    font-family: 'Mirza';
  }

</style>