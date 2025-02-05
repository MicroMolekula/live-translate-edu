<script setup>
import {
  Room,
  RoomEvent,
  Track,
} from 'livekit-client';
import {ref} from "vue";

const result = ref('')
const token1 = 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3Mzg4NTk4ODEsImlzcyI6ImRldmtleSIsIm5iZiI6MTczODc3MzQ4MSwic3ViIjoiaXZhbiIsInZpZGVvIjp7InJvb20iOiJteXJvb20iLCJyb29tSm9pbiI6dHJ1ZX19.USfSLOQkYzzzYuokE1F6pvoh8_sGagV0GxQropbaFYI'

const room = new Room()

async function connectionRTC() {
  let token = token1
  const decoder = new TextDecoder()

  await room.prepareConnection('http://localhost:7880', token);
  await fetch("http://localhost:8080/api/connect")
  room
    .on(RoomEvent.TrackSubscribed, handleTrackSubscribed)
      .on(RoomEvent.DataReceived, function (payload, participant, kind) {
        result.value = decoder.decode(payload)
      })
  await room.connect('ws://localhost:7880', token);
  console.log('connected to room', room.activeSpeakers);
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
  await fetch("http://localhost:8080/api/disconnect")
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
</script>

<template>
  <div class="flex flex-col h-screen">
    <!-- Header -->
    <header class="bg-green-600 text-white p-4 flex justify-between items-center">
      <div class="text-2xl font-semibold">Перевод в реальном времени</div>
      <div class="text-lg">Красиков Иван</div>
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
          Подключиться
        </button>
        <button @click="disconnectRecognize" class="mt-4 px-6 py-2 bg-red-500 text-white rounded-lg hover:bg-red-600 focus:outline-none focus:ring-2 focus:ring-blue-500">
          Отключиться
        </button>
      </div>
    </div>
  </div>
</template>

<style scoped>

</style>