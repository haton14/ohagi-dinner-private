<script lang="ts">
	import Day from './Day.svelte';
	import FoodRow from './FoodRow.svelte';
</script>

<div class="w-4/5 flex flex-row">
	<Day unixMicroTime={12345} />
	<div class="w-4/5 flex flex-col">
		<FoodRow foodName={'ハムセレ'} foodQuantity={6.7} foodUnit={'g'} />
		<FoodRow foodName={'ハムセレ'} foodQuantity={9.8} foodUnit={'g'} />
	</div>
</div>
