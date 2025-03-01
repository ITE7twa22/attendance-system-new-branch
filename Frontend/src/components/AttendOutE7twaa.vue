<template>
  <div>
    <div class="loader-container" id="loadingScreen">
      <div class="gear">
        <img src="../assets/img/Loader_Image.png" alt="Loading" />
      </div>
    </div>

    <div class="flex flex-col items-center justify-center mb-6">
      <div class=" bg-[#282828e4] p-8 rounded-lg shadow-lg w-full max-w-md opacity-90 hover:opacity-100 relative mb-4">
        <p class="text-gray-400 flex items-center justify-start flex-row-reverse mb-4">
          <img src="../assets/img/cautionIcon.png" class="w-4 h-4 mr-1 ml-1" alt="Caution Icon"  @click="alert('في حال واجهتك أي مشكلة بالتسجيل، الرجاء التوجه لفريق الموارد البشرية')"/>
          <span>
            أدخل<span class="text-e7twa22-gray"> رقم التسجيل</span> لداخل المؤسسة
            <span class="text-e7twa22-gray">أو رقم الهوية</span> لخارج المؤسسة
          </span>
        </p>

        <input
          @keyup.enter.prevent="getRelationsInfo"
          @input="handleInputChange"
          v-model="volunteerID"
          type="text"
          placeholder=" أدخل رقم التسجيل أو الهوية هنا"
          class="border border-gray-300 border-dashed px-4 py-2 rounded-lg w-full shadow-md shadow-[#5daaae49] focus:outline-none focus:ring-2 focus:ring-[#5daaae] text-center"
          inputmode="numeric"
        />
        <p v-if="inputError" class="text-red-500" style="text-align: center;">{{ inputError }}</p>
        <div class="flex justify-center mt-4 w-full">
          <button
    @click="getRelationsInfo"
    v-if="isInputChanged && isSearch"
    class="btn text-white px-4 py-2 rounded-md cursor-pointer"
    style="border-radius: 12px;">
    التالي
    <img src="../assets/img/NextIcon.png" alt="Next Icon" class="inline-block ml-2" style="width: 10px; height: 10px;" />
  </button>

        </div>
     
      </div >
      <p v-if="isFindName" class="text-e7twa22-white 4xl">
  أهلًا بك <span class="text-e7twa22-blue">{{ relationsVolunteerName }}</span> هل تريد تأكيد تسجيل الدخول؟
</p>
<br />
<p v-if="isHaveCode && isFindName"  class="text-2xl text-white font-light">رقم التسجيل الخاص بك هو <span class="text-e7twa22-blue font-normal">{{ code }}</span></p>
<br />

<div v-if="isFindName" class="button-container flex justify-center gap-4 mt-4">

  <button @click="confirmLogin" style="background-color: rgba(93, 170, 174, 0.8);" class="text-white px-4 py-2 rounded-md cursor-pointer">دخول</button>

  <button @click="cancelLogin" class="bg-gray-500 text-white px-4 py-2 rounded-md cursor-pointer">إلغاء</button>

</div>

<p v-if="ifSucsLogin" class="text-e7twa22-white 4xl">
  أهلًا بك <span class="text-e7twa22-blue">{{ relationsVolunteerName }}</span> هل تريد تأكيد تسجيل الخروج؟
</p>
<br />
<p v-if="isHaveCode && ifSucsLogin"  class="text-2xl text-white font-light">رقم التسجيل الخاص بك هو <span class="text-e7twa22-blue font-normal">{{ code }}</span></p>

<br />

<div v-if="ifSucsLogin" class="button-container flex justify-center gap-4 mt-4">

  <button @click="confirmLogout" class="bg-red-500 text-white px-4 py-2 rounded-md cursor-pointer">خروج</button>


  <button @click="cancelLogout" class="bg-gray-500 text-white px-4 py-2 rounded-md cursor-pointer">إلغاء</button>
 

</div>
<div v-if="isSuccessLoginPopup" class="popup">
        <span class="close-btn" @click="isSuccessLoginPopup = false">&times;</span>
        <p class="text-center text-lg">تم تسجيل دخولك بنجاح!</p>
        <button @click="isSuccessLoginPopup = false" class="btn">إغلاق</button>
      </div>
<div v-if="isLogoutPopupVisible" class="popup">
  <span class="close-btn" @click="isLogoutPopupVisible = false">&times;</span>
  <p class="text-center text-lg">تم تسجيل خروجك بنجاح!</p>
  <button @click="isLogoutPopupVisible = false" class="btn">إغلاق</button>
</div>
     

<!-- Modal for Success/Error Messages -->

    <!-- رسائل النجاح/الخطأ في الـ messageModal -->
    <div
      id="messageModal"
      tabindex="-1"
      aria-hidden="true"
      :class="{'hidden': !modalVisible, 'flex': modalVisible}"
      class="fixed inset-0 z-50 flex justify-center items-center bg-black bg-opacity-50 transition-opacity duration-300"
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
                d="M4.293 4.293a1 1 0 011.414 0L10 8.586l4.293-4.293a1 1 0 111.414 1.414L11.414 10l7.293 7.293a1 1 0 01-1.414 1.414L10 11.414l-4.293 4.293a1 1 0 01-1.414-1.414L8.586 10 4.293 5.707a1 1 0 010-1.414z"
                clip-rule="evenodd"
              ></path>
            </svg>
            <span class="sr-only">إغلاق النافذة</span>
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
                d="M16.707 5.293a1 1 0 010 1.414L10 14.414l-6.707-6.707a1 1 0 111.414-1.414L10 11.586l5.293-5.293a1 1 0 011.414 0z"
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

<div class="text-center flex flex-col gap-4">
  <p v-if="catchName" class="text-4xl text-white font-light">{{ catchName }}</p>
</div>


    </div>

    <br />
    <br />

 



    <div v-if="isAddVolunteerFormVisible" class="popup">
  <span class="close-btn" @click="closeAddVolunteerForm">&times;</span>

  <!-- Form fields for adding a new volunteer -->
  <form @submit.prevent="submitAddVolunteerForm" class="form-container">
    <div class="form-group">
      <div class="input-group">
        <label for="newVolunteerName" dir="rtl" class="center-label">:اسم المتطوع / ـة</label>
        <input v-model="newVolunteerName" type="text" id="newVolunteerName" name="newVolunteerName" required dir="rtl" class="border-input" @input="validateName" />
        <p v-if="nameError" class="text-red-500">{{ nameError }}</p>
      </div>

      <div class="input-group">
        <label for="newVolunteerId" class="center-label">:رقم الهوية</label>
        <input v-model="newVolunteerId" type="text" id="newVolunteerId" name="newVolunteerId" maxlength="10" required dir="rtl" class="border-input" />
        <p v-if="idError" class="text-red-500">{{ idError }}</p>
      </div>

      <div class="input-group">
        <label for="newVolunteerPhone" class="center-label">:رقم الجوال</label>
        <input v-model="newVolunteerPhone" type="text" id="newVolunteerPhone" name="newVolunteerPhone" required dir="rtl" class="border-input" />
        <p v-if="phoneError" class="text-red-500">{{ phoneError }}</p>
      </div>

      <div class="input-group">
        <label for="newVolunteerEmail" class="center-label">:البريد الإلكتروني</label>
        <input v-model="newVolunteerEmail" type="email" id="newVolunteerEmail" name="newVolunteerEmail" required dir="rtl" class="border-input" @input="validateEmail" />
        <p v-if="emailError" class="text-red-500">{{ emailError }}</p>
      </div>

      <div class="input-group">
        <label for="newVolunteerGender" class="center-label">:الجنس</label>
        <div class="radio-group">
          <input type="radio" id="male" value="male" v-model="newVolunteerGender" name="gender">
          <label for="male">متطوع</label>
          <input type="radio" id="female" value="female" v-model="newVolunteerGender" name="gender" checked>
          <label for="female">متطوعة</label>
        </div>
      </div>
    </div>
    
    <p class="text-center flex flex-col gap-4" style="color: red;">{{ addvError }}</p>
    <button type="submit" class="add-button" style="background-color: #5DAAAD;" :disabled="emailError">إضافة</button>
  </form>
</div>


<div
  id="successPopupModal"
  tabindex="-1"
  aria-hidden="true"
  :class="{'hidden': !isSuccessPopupVisible, 'flex': isSuccessPopupVisible}"
  class="fixed inset-0 z-50 flex justify-center items-center bg-black bg-opacity-50 transition-opacity duration-300"
>
  <div class="relative p-4 w-full max-w-xs mx-auto">
    <div
      class="relative p-4 text-center rounded-lg shadow-lg bg-white"
    >
      <button
        id="closeSuccessPopupModal"
        @click="isSuccessPopupVisible = false"
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
            d="M4.293 4.293a1 1 0 011.414 0L10 8.586l4.293-4.293a1 1 0 111.414 1.414L11.414 10l7.293 7.293a1 1 0 01-1.414 1.414L10 11.414l-4.293 4.293a1 1 0 01-1.414-1.414L8.586 10 4.293 5.707a1 1 0 010-1.414z"
            clip-rule="evenodd"
          ></path>
        </svg>
        <span class="sr-only">إغلاق النافذة</span>
      </button>
      <div
        id="modalIcon"
        class="w-12 h-12 rounded-full p-1 flex items-center justify-center mx-auto mb-3.5 bg-green-200"
      >
        <svg
          id="modalSvg"
          aria-hidden="true"
          class="w-8 h-8 text-green-500"
          fill="currentColor"
          viewBox="0 0 20 20"
          xmlns="http://www.w3.org/2000/svg"
        >
          <path
            id="successPath"
            fill-rule="evenodd"
            d="M16.707 5.293a1 1 0 010 1.414L10 14.414l-6.707-6.707a1 1 0 111.414-1.414L10 11.586l5.293-5.293a1 1 0 011.414 0z"
            clip-rule="evenodd"
          ></path>
        </svg>
      </div>
      <p
        id="modalMessage"
        class="mb-2 text-2xl font-semibold text-center text-gray-900"
      >
        تم إضافة المتطوع بنجاح! الرجاء تسجيل الدخول.
      </p>
    </div>
  </div>
</div>
     

<div v-if="isVolunteerRegistrationVisible" class="text-center">
  <a
    @click="openAddVolunteerForm"
    class="text-white px-4 py-2 rounded-md mt-4 cursor-pointer"
    style="font-size: 1.5rem; margin-top: -40px;"
  >
    في حال كنت متطوع جديد يُرجى تسجيل بياناتك <span class="text-e7twa22-blue font-normal"> بالضغط هنا </span>
  </a>
</div>
  </div>
</template>

<script>
import axios from 'axios';



export default {
  data() {
    return {
      addvError: "",
      isHaveCode: false,
      enableButton: false,
      isSearch: true,
      isVolunteerRegistrationVisible: false,
      code: null,
      volunteerID: '',
      relationsVolunteerName: '',
      catchName: '',
      isLogin: false,
      isFindName: false,
      ifCheckIn: false,
      checkinOutflag: false,
      ifSucLogin: false,
      ifSucsLogin: false,
      isAddVolunteerFormVisible: false,
      newVolunteerName: '',
      newVolunteerId: '',
      newVolunteerGender: 'female',
      newVolunteerPhone: '',
      nameError: "",
      idError: "",
      phoneError: "",
      isSuccessPopupVisible: false,
      isSuccessLoginPopup:false,
      transitionToLogin: false,
      newVolunteerEmail: '',
      emailError: '',
      inputError: "",
      isLogoutPopupVisible: false,
      modalMessage: '',
    modalVisible: false,
    isSuccess: false,
    };
  },
  computed: {


  },

  methods: {

    showMessageModal(isSuccess, message) {
    this.modalMessage = message;
    this.modalVisible = true;
    this.isSuccess = isSuccess;

 
    setTimeout(() => {
      this.closeModal();
    }, 1500);
  },

    validateEmail() {
      const emailPattern = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
      if (!emailPattern.test(this.newVolunteerEmail)) {
        this.emailError = 'البريد الإلكتروني غير صحيح، يرجى إدخال بريد صالح مثل example@gmail.com';
      } else {
        this.emailError = '';
      }
    },

    validateName() {
      const nameParts = this.newVolunteerName.trim().split(" ");
      if (nameParts.length < 3) {
        this.nameError = 'يجب إدخال الاسم الثلاثي';
      } else {
        this.nameError = '';
      }
    },

  closeModal() {
    this.modalVisible = false;
    this.isSuccess = false;
  },
    confirmLogin() {
    this.showLoadingScreen();

    this.handleCheckIn();
    this.showMessageModal(true, "تم تسجيل دخولك بنجاح!");
    this.isHaveCode = false;
    this.volunteerID = '';
    this.code = null;

    setTimeout(() => {
      this.isSuccessLoginPopup = false;
    }, 2000);
  },

  cancelLogin() {
 
    console.log("تسجيل الدخول ملغي");
    this.isFindName = false;
  this.volunteerID = '';
  this.relationsVolunteerName = '';
  this.isSuccessLoginPopup = false;
  this.modalVisible = false;
    // window.location.reload();
    this.isFindName = false;
    this.volunteerID = '';
  },

  confirmLogout() {
    this.showLoadingScreen();
     this.handleCheckIn();
     this.showMessageModal(true, "تم تسجيل خروجك بنجاح!");
     console.log("ifCheckIn:", this.ifCheckIn);
    this.isHaveCode = false;
    this.volunteerID = '';
this.code = null;
this.ifCheckIn = false;
    setTimeout(() => {
      this.isLogoutPopupVisible = false;
    }, 2000);
  },

  cancelLogout() {
    console.log("تسجيل الخروج ملغي");
    this.ifSucsLogin = false;
  this.volunteerID = '';
  this.relationsVolunteerName = '';
  this.isLogoutPopupVisible = false;
  this.modalVisible = false;
    this.isFindName = false;
    this.volunteerID = '';
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
    submitAddVolunteerForm() {
  // إعادة تعيين رسائل الأخطاء
  this.nameError = "";
  this.idError = "";
  this.phoneError = "";
  this.addvError = "";
 

  let isValid = true;

  // التحقق من الاسم
  if (!this.newVolunteerName || this.newVolunteerName.length < 4 || /[^ا-يء\s]/.test(this.newVolunteerName)) {
    this.nameError = "يجب أن يحتوي الاسم على 4 أحرف عربية على الأقل ولا يحتوي على أرقام أو رموز.";
    isValid = false;
  }

  // التحقق من رقم الهوية
  if (!/^\d{10}$/.test(this.newVolunteerId)) {
    this.idError = "رقم الهوية يجب أن يكون 10 أرقام فقط.";
    isValid = false;
  }

  // التحقق من رقم الجوال
  if (!/^05\d{8}$/.test(this.newVolunteerPhone)) {
    this.phoneError = "رقم الجوال يجب أن يبدأ بـ 05 وأن يكون 10 أرقام فقط.";
    isValid = false;
  }

  if (isValid) {
    this.showLoadingScreen(); // Show loading before the API call

    const formData = {
      Name: this.newVolunteerName,
      NationalId: parseInt(this.newVolunteerId),
      Phone: this.newVolunteerPhone,
      Gender: this.newVolunteerGender,
    };

    console.log(formData);

    axios.post('http://localhost:8080/addNewVolunteer', JSON.stringify(formData), {
      headers: {
        'Content-Type': 'application/json',
      },
    })
    .then(response => {
      this.addvError = "";
      // إعادة تعيين الحقول
      this.newVolunteerName = '';
      this.newVolunteerPhone = '';
      this.newVolunteerGender = 'female';
      this.isAddVolunteerFormVisible = false;
      this.addvError = "";

      this.isSuccessPopupVisible = true;


      setTimeout(() => {
        this.isSuccessPopupVisible = false;
        this.volunteerID = this.newVolunteerId;
        this.getRelationsInfo();
        this.transitionToLogin = true;
      }, 2000);
   
    })
    .catch(error => {
      this.addvError = "رقم الهوية مسجل مسبقا";
      console.error('Error adding volunteer:', error);
    })
    .finally(() => {
      this.hideLoadingScreen(); // Hide loading after the API call
    });
  } else {
   
    this.hideLoadingScreen();
  }
},

handleInputChange() {
  // إعادة تعيين الرسائل السابقة
  this.catchName = "";
  this.addvError = "";
  this.enableButton = false;
  this.isSearch = true;
  this.ifSucsLogin = false;
  this.isFindName = false;
  this.ifSucLogin = false;
  this.catchName = '';
  this.isLogin = false;
  this.isHaveCode = false;
  this.relationsVolunteerName = "";
  this.ifCheckIn = false;
  this.isInputChanged = true;

 
  // التحقق من رقم الهوية (أقل من 10 أو أكثر من 10 أرقام)
  if (this.volunteerID.length > 5 && (this.volunteerID.length < 10 || this.volunteerID.length > 10)) {
    this.inputError = "رقم الهوية يجب أن يتكون من 10 أرقام فقط";
    this.enableButton = false;
    return;
  }

  // إذا كانت المدخلات فارغة، إزالة رسالة الخطأ
  else if (this.volunteerID.length === 0) {
    this.inputError = "";
    
  }

  // إذا كانت المدخلات صحيحة (رقم التسجيل 5 أرقام أو أكثر ورقم الهوية 10 أرقام)
  else if (this.volunteerID.length === 10 || this.volunteerID.length <= 5 ) {
    this.inputError = "";
    this.enableButton = true;
  }

  // الاستمرار في استرجاع المعلومات
  this.getIfCheckIn();
},


    handleCheckIn() {
   

      this.showLoadingScreen(); // Show loading before the API call
      console.log("الرقم المستخدم:", this.code || this.volunteerID); // تحقق من الرقم المستخدم
      const endpoint = this.ifCheckIn ? 'http://localhost:8080/updateLogoutDateTime' : 'http://localhost:8080/addRelationsVolunteer';
      const loginDateTime = new Date();
        // استخدام الكود إذا كان موجودًا
  const idToUse = this.code ? parseInt(this.code) : parseInt(this.volunteerID);
  console.log("الرقم المستخدم لتسجيل الدخول:", idToUse);

      // if (this.code) {
       
       axios.post(endpoint, {
        volunteerID: idToUse,
       
         RelationsVolunteerName: this.relationsVolunteerName,
         LoginDateTime: loginDateTime.toISOString(),
     
      })
      .then(response => {
        this.enableButton = false;
        this.isSearch = true;

        if (this.ifCheckIn) {
          this.ifCheckIn = false;
          this.isLogin = true;
          this.ifSucsLogin = false;
          this.ifSucLogin = false;
        } else {
          this.ifCheckIn = true;
          this.ifSucsLogin = false;
          this.ifSucLogin = true;
          this.isLogin = true;
        }

        this.volunteerID = '';
        this.isFindName = false;
      })
      .catch(error => {
        console.log("ERROR!!!!", error);
      })
      .finally(() => {
        this.hideLoadingScreen(); // Hide loading after the API call
      });
 
},

    async   getRelationsInfo() {

      this.inputError = "";
      console.log("رقم التسجيل:", this.volunteerID);
      const numericPattern = /^\d+$/;
      if (!numericPattern.test(this.volunteerID)) {
    this.inputError = "يرجى إدخال أرقام فقط";
    return;
  }
     this.showLoadingScreen(); //  this.showLoadingScreen();
      axios
        .post(
          'http://localhost:8080/getNumberHandler',
          { volunteerID: parseInt(this.volunteerID) },
          {
            headers: {
              'Content-Type': 'application/json',
            },
          }
        )
        .then(async (response) => {
          console.log('Response Data:', response.data); //  (داخل المؤسسة  VolunteerID موجود هنا

          if (response.data.VolunteerID) {
            this.relationsVolunteerName = response.data.RelationsVolunteerName;
            this.isVolunteerRegistrationVisible = false;
            this.code = response.data.VolunteerID;

            console.log("Response Data:", response.data);
            this.enableButton = true
            this.isSearch = false
            this.catchName = '';
            this.isHaveCode = true;
            this.relationsVolunteerName = response.data.RelationsVolunteerName;
            this.isFindName = true;
            this.isVolunteerRegistrationVisible = false;
            console.log("قيمة code قبل الاستدعاء:", this.code);
            this.code = response.data.VolunteerID;
            console.log("VolunteerID من الاستجابة:", response.data.VolunteerID);
            await   this.getIfCheckIn(this.code);
            console.log("Volunteer ID:", this.volunteerID);
            console.log("Relations Volunteer Name:", this.relationsVolunteerName);
            console.log("Is Find Name:", this.isFindName);
            console.log("If Successful Login:", this.ifSucsLogin);
            if (this.ifCheckIn != true) {
              console.log(this.ifCheckIn);

              this.isFindName = true
            }
            else {
              console.log(this.ifCheckIn);

              this.ifSucsLogin = true
            }
          }
          else {
                    // يوجد من خارج المؤسسة
            console.log("Response Data:", response.data);
            this.enableButton = true
            this.isSearch = false
            this.catchName = '';
            this.isHaveCode = false
            this.relationsVolunteerName = response.data.RelationsVolunteerName;
            this.isVolunteerRegistrationVisible = false;
            console.log("VolunteerID من الاستجابة:", response.data.VolunteerID);
            await   this.getIfCheckIn(this.volunteerID);
            if (this.ifCheckIn != true) {

              this.isFindName = true
            }
            else {
              this.ifSucsLogin = true
            }
          }


        })

        .catch((error) => {
          this.relationsVolunteerName = '';
          this.isFindName = false;
          this.isHaveCode = false
          this.isVolunteerRegistrationVisible = true;
         
          
          if (this.volunteerID.length < 6 ) {
    this.catchName = 'الكود غير صالح، يرجى تسجيل الدخول باستخدام رقم الهوية';
    
    return;
  }
          else if (this.volunteerID.length === 10 ) {
    this.catchName = 'رقم الهوية غير موجود';
    
    return;
  }
  

        })
        .finally(() => {
        this.hideLoadingScreen(); // Hide loading after the API call
      });
    },
    async getIfCheckIn(id) {



      try {


    // تحقق من وجود المتطوع في سجلات الحضور باستخدام volunteerID
    console.log(id);
    const attendanceResponse = await axios.post('http://localhost:8080/searchAttendanceHandler', { volunteerID: parseInt(id) });
    console.log("استجابة الخادم:", attendanceResponse.data);
    if (attendanceResponse.data === 'Volunteer not came today') {
      this.ifCheckIn = false;
    } else if (attendanceResponse.data === 'Volunteer came today') {
      this.isFindName = false
      this.ifCheckIn = true;
    }
  } catch (error) {
    console.error('خطأ أثناء التحقق من حالة تسجيل الدخول:', error);
  }
    },
    openAddVolunteerForm() {
      this.isAddVolunteerFormVisible = true;
      this.isHaveCode = false;
      this.newVolunteerId = this.volunteerID;
    },
    closeAddVolunteerForm() {
      this.isAddVolunteerFormVisible = false;
      this.addvError = "";
      this.newVolunteerName = '';
      this.newVolunteerId = '';
      this.newVolunteerPhone = '';
    },
  },
};
</script>


<style>
 @import url('https://fonts.googleapis.com/css2?family=Mirza:wght@400;500;600;700&display=swap');

.btn {
  background-color: rgb(93, 170, 174);
  color: white;
  border-radius: 20px;
  cursor: pointer;
  transition: background-color 0.3s ease;
}

.text-e7twa22-blue {
  color: rgba(93, 170, 174, 0.8);
  font-weight: 500;
  font-family: 'Mirza';
}
.text-e7twa22-white {
  color: rgba(255, 255, 255, 0.8);
  font-weight: 500;
  font-family: 'Mirza';
  font-size: 2.25rem;
  text-align: center;
}

.text-e7twa22-gray {
  color: #efede9;
  font-weight: bold;
  font-family: 'Mirza';
}

.loader-container {
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

.popup {
  position: fixed;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  background-color: white;
  padding: 20px;
  border-radius: 8px;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.2);
  z-index: 1000;
}

.close-btn {
  position: absolute;
  top: 10px;
  right: 10px;
  cursor: pointer;
}

.form-container {
  display: flex;
  flex-direction: column;
}

.form-group {
  display: flex;
  flex-direction: column;
  align-items: flex-end;
  margin-bottom: 10px;
}

.input-group {
  display: flex;
  flex-direction: column;
  align-items: flex-end;
  margin-bottom: 10px;
  text-align: right !important;;
}

.input-group label {
  margin-bottom: 5px;
}

.center-label {
  text-align: right;
  /* text-align: center; */
}

.border-input {
  border: 1px solid #ccc;
  padding: 8px;
  border-radius: 4px;
}

.radio-group {
  display: flex;
  flex-direction: row;
  align-items: center;
}

.radio-group input {
  margin-left: 10px;
}

.radio-group label {
  margin-left: 5px;
}
.add-button {
  background-color: #5DAAAD;
  color: white;
  padding: 10px 20px;
  border-radius: 8px;
  font-size: 16px;
  font-weight: bold;
  cursor: pointer;
  transition: background-color 0.3s ease;
  box-shadow: 0px 4px 6px rgba(0, 0, 0, 0.1);
}
.label {
 
    text-align: right;
    display: block;  
}
.add-button:hover {
  background-color: #4a8f8f;
}
</style>