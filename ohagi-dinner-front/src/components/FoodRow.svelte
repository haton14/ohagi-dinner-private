<script lang="ts">
	export let foodName: string;
	export let foodQuantity: number;
	export let foodUnit: string;
</script>

<div class="flex flex-row justify-between">
	<div class="w-4/5">{foodName}</div>
	<div class="w-1/5 flex flex-row justify-between">
        <div>{foodQuantity}</div>
		<div>{foodUnit}</div>
	</div>
</div>
