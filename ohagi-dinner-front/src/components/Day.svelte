<script lang="ts">
	export let unixMicroTime: number;
	let date = new Date(unixMicroTime);
</script>

<div class="w-1/5 flex flex-col text-center">
	<div>火</div>
	<div>{date.getDate()}</div>
	<div>{String(date.getHours()).padStart(2, '0')}:{String(date.getMinutes()).padStart(2, '0')}</div>
</div>
