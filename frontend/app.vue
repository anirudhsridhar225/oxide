<template>
  <div>
		<input
			v-model="name"
			type="text"
			placeholder="Enter your name"
			class="border px-4 py-2 rounded-lg"
		/>

		<button
			@click="fetchGreeting"
			class="bg-blue-500 text-white px-4 py-2 rounded hover:bg-blue-600"
		>
			Greet me
		</button>

		<div
			v-if="greeting"
			class="mt-4 p-2 border rounded"
		>
			{{greeting}}
		</div>
		
  </div>
</template>

<script setup>
import {ref} from "vue";

const name = ref('')
const greeting = ref('')

const fetchGreeting = async () => {
  if (!name.value) return

  try {
    const { data } = await useFetch(`http://localhost:8000/api/greet/${name.value}`, {
      method: 'GET'
    })

    greeting.value = data.value.message
  } catch (err) {
    console.error('Error fetching greeting:', err)
    greeting.value = 'Failed to fetch greeting.'
  }
}
</script>
