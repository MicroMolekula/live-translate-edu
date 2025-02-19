<script setup>
import {
  Room,
  RoomEvent,
  Track,
} from 'livekit-client';
import {onMounted, ref} from "vue";
import {useRouter} from "vue-router";

const result = ref('')
const token1 = localStorage.getItem("room_token")

const room = new Room()
const router = useRouter()
const user = ref({
  name: "",
  surname: "",
})

async function getUser() {
  let response = await fetch('http://localhost:8080/api/me', {
    headers: {
      'Authorization': 'Bearer ' + localStorage.getItem('jwt')
    },
  })
  let result = await response.json()
  if (response.status === 200) {
    user.value = result
  }
}

onMounted(async function () {
  await getUser()
  await room.prepareConnection('http://localhost:7880', token1);
  await room.connect('ws://localhost:7880', token1);
  room.on(RoomEvent.DataReceived, function (payload, participant, kind) {
    result.value = decoder.decode(payload)
  })
})

const decoder = new TextDecoder()



async function connectionRTC() {
  await fetch("http://localhost:8080/api/connect", {
    headers: {
      'Authorization': 'Bearer ' + localStorage.getItem('jwt')
    }
  })
  room
    .on(RoomEvent.TrackSubscribed, handleTrackSubscribed)
  navigator.mediaDevices.getUserMedia({
    audio: {
      sampleRate: 48000
    }
  }).then((stream) => {
    const audioTrack = stream.getAudioTracks()[0];
    room.localParticipant.publishTrack(audioTrack);
  }).catch((error) => {
    console.error("Ошибка при доступе к микрофону:", error);
  });
}

async function disconnectRecognize() {
  await room.disconnect()
  await fetch("http://localhost:8080/api/disconnect", {
    headers: {
      'Authorization': 'Bearer ' + localStorage.getItem('jwt')
    }
  })
}

function handleTrackSubscribed(
    track,
    publication,
    participant,
) {
  if (track.kind === Track.Kind.Video || track.kind === Track.Kind.Audio) {
    // attach it to a new HTMLVideoElement or HTMLAudioElement
    attachTrack(track, participant)
  }
}

function attachTrack(track, participant) {
  const v = document.getElementById("remoteAudio");
  track.attach(v);
}

function logout() {
  localStorage.clear()
  router.push("/auth")
}
</script>

<template>
  <div class="flex flex-col h-screen">
    <!-- Header -->
    <header class="bg-green-600 text-white p-4 flex justify-between items-center">
      <div class="text-2xl font-semibold">Перевод в реальном времени</div>
      <div class="flex flex-row gap-3">
        <div class="text-lg pt-5">{{user.name + " " + user.surname}}</div>
        <button @click="logout" class="mt-4 px-6 py-2 bg-red-500 text-white rounded-lg hover:bg-red-600 focus:outline-none focus:ring-2 focus:ring-blue-500">
          Выйти
        </button>
      </div>
    </header>

    <!-- Main Content -->
    <div class="flex flex-col flex-1">
      <!-- Subtitles Section -->
      <div class="flex-1 flex items-center justify-center p-6 flex-col gap-1">
        <div class="bg-gray-800 text-white text-xl p-6 rounded-lg shadow-md w-full max-w-3xl">
          <p v-if="result" class="whitespace-pre-wrap">{{ result }}</p>
          <p v-else class="text-center text-gray-400">Место для субтитров</p>
        </div>
        <!-- Button under subtitles -->
        <button @click="connectionRTC" class="mt-4 px-6 py-2 bg-green-500 text-white rounded-lg hover:bg-green-600 focus:outline-none focus:ring-2 focus:ring-blue-500">
          Запустить перевод речи
        </button>
        <button @click="disconnectRecognize" class="mt-4 px-6 py-2 bg-red-500 text-white rounded-lg hover:bg-red-600 focus:outline-none focus:ring-2 focus:ring-blue-500">
          Отключить перевод речи
        </button>

      </div>
    </div>
  </div>
</template>

<style scoped>

</style>