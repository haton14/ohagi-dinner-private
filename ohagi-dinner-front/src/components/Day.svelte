<script lang="ts">
	export let unixMicroTime: number;
	let date = new Date(unixMicroTime);
	let days = ['月', '火', '水', '木', '金', '土', '日'];
</script>

<div class="w-1/5 flex flex-col justify-around text-center">
	<p class="text-xs text-gray-700">{days[date.getDay()]}</p>
	<p class="text-xl">{date.getDate()}</p>
	<p class="text-xs text-gray-400">{String(date.getHours()).padStart(2, '0')}:{String(date.getMinutes()).padStart(2, '0')}</p>
</div>
