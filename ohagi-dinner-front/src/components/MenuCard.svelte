<script lang="ts">
	import Day from './Day.svelte';
	import FoodRow from './FoodRow.svelte';
</script>

<div class="flex flex-row p-1 px-6">
	<Day unixMicroTime={12345} />
	<div class="w-4/5 flex flex-col justify-around border border-grey-light rounded">
		<FoodRow foodName={'ハムセレ'} foodQuantity={6.7} foodUnit={'g'} />
		<FoodRow foodName={'ハムセレ'} foodQuantity={9.8} foodUnit={'g'} />
	</div>
</div>
