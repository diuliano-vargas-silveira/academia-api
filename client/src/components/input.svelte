<script lang="ts">
    export let label: string
    export let name: string
    export let value: string = ''
    export let mask: string = ''
    export let formatter: ((value: string, mask: string) => string) | null
    export let validator: ((value: string) => string) | null

    $: value = onChange(value)


    function onChange(value: string): string {
        const regex: RegExp = /[^0-9]/g

        value = value.replaceAll(regex, '').trim() 
        
        

        return formatter(value, mask)
    }
    
</script>

<div>
    {`${label}:${name}`}

    <input 
        bind:value
        />
</div>

<style>

</style>