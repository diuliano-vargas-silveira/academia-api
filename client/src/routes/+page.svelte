<script lang="ts">
    import Input  from "../components/input.svelte"
    import { InputUtils } from "$lib/utils/input.utils"

    const inputUtils = new InputUtils()
    let value = ''


    function validator(input: string):string {
        return input
    }

    function format(input: string, mask: string): string {
        return inputUtils.format(input, mask)
    }

</script>


<h1>Welcome to SvelteKit</h1>
<p>Visit <a href="https://kit.svelte.dev">kit.svelte.dev</a> to read the documentation</p>
<Input name={'nome'} label={'label'} mask='XXX.XXX.XXX-XX' value={value} formatter={inputUtils.format} validator={validator} /> 
