<script setup>
import ButtomPanel from "@/components/lesson/ButtomPanel.vue";
import SideBarChat from "@/components/lesson/SideBarChat.vue";
import {onMounted, ref} from "vue";
import Subtitle from "@/components/lesson/Subtitle.vue";
import {roomStore, userStore} from "@/stores/stores.js";
import {onBeforeRouteLeave, useRoute, useRouter} from "vue-router";
import {
  Room,
  RoomEvent,
  Track,
} from 'livekit-client';
import {connectRoom, disconnectRoom, getUsersInLesson} from "@/lib/request/lesson.js";
import SideBarUsers from "@/components/lesson/SideBarUsers.vue";

const decoder = new TextDecoder()

const chatOpen = ref(false)
const usersOpen = ref(false)
const userData = userStore()
const roomData = roomStore()
const room = new Room()
const unmuteMicro = ref(false)
const usersConnects = ref([])

const subtitle = ref('')
const wsClient = ref()

const route = useRoute()
const router = useRouter()

onMounted(async () => {
  await room.prepareConnection('http://localhost:7880', roomData.mapRoomToken.roomToken);
  await room.connect('ws://localhost:7880', roomData.mapRoomToken.roomToken);
  wsClient.value =
  room.on(RoomEvent.DataReceived, function (payload, participant, kind) {
    subtitle.value = decoder.decode(payload)
  })
  if (userData.user.role === 'teacher') {
    await connectionRTC()
  }
})

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

async function connectionRTC() {
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

onBeforeRouteLeave((to, from, next) => {
  room.disconnect()
  if (unmuteMicro.value) {
    disconnectRoom(userData.token, roomData.mapRoomToken.roomName)
        .then((result) => {
          console.log('Выключение распознавания речи')
        })
        .catch((e) => {
          console.log('Ошибка выключения распознавания речи')
        })
  }
  next()
})

function mute(event){
  const audioTrackPublications = room.localParticipant.audioTrackPublications;
  unmuteMicro.value = event
  if (unmuteMicro.value) {
    for (const publication of audioTrackPublications.values()) {
      const audioTrack = publication.track;
      if (audioTrack && audioTrack.kind === Track.Kind.Audio) {
        audioTrack.unmute();
        console.log('Microphone unmuted');
      }
    }
    connectRoom(userData.token, roomData.mapRoomToken.roomName)
        .then((result) => {
          console.log('Включение распознавания речи')
        })
        .catch((e) => {
          console.log('Ошибка включения распознавания речи')
        })
  } else {
    const audioTrackPublications = room.localParticipant.audioTrackPublications;
    for (const publication of audioTrackPublications.values()) {
      const audioTrack = publication.track;
      if (audioTrack && audioTrack.kind === Track.Kind.Audio) {
        audioTrack.mute();
        console.log('Microphone muted');
      }
    }
    disconnectRoom(userData.token, roomData.mapRoomToken.roomName)
        .then((result) => {
          console.log('Выключение распознавания речи')
        })
        .catch((e) => {
          console.log('Ошибка выключения распознавания речи')
        })
  }
}

function leave() {
  router.push('/')
}

function openChat(msg) {
  if (usersOpen.value) {
    usersOpen.value = false
  }
  chatOpen.value = !chatOpen.value
}

function openUsers() {
  if (chatOpen.value) {
    chatOpen.value = false
  }
  usersOpen.value = !usersOpen.value
  if (usersOpen.value) {
    getUsersInLesson(userData.token, roomData.mapRoomToken.roomName)
        .then((result) => {
          console.log(result.data.users)
          usersConnects.value = result.data.users
        })
        .catch((e) => {
          console.log(e.message)
        })
  }
}
</script>

<template>
  <div class="h-screen bg-main">
    <div class="flex flex-row gap-2 bg-main text-white" style="height: calc(100vh - 140px)">
      <div class="flex-1 grid gap-4 p-4 place-items-center">
        <Subtitle :text="subtitle"></Subtitle>
      </div>
      <SideBarChat :open="chatOpen"/>
      <SideBarUsers :users-chat="usersConnects" :open="usersOpen"/>
    </div>
    <ButtomPanel class="h-20" style="min-height: 80px" @chat="openChat" @users="openUsers" @mute="mute" @leave="leave"></ButtomPanel>
  </div>
</template>

<style scoped>
.bg-main {
  background-color: rgba(32, 33, 36, 1.00);
  color: white;
}
</style>