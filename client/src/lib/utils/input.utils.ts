export class InputUtils {

    replacer: string = 'X'

    format(input: string, mask: string): string {
        let formattedString = '' 
        let indexInput = 0
    
        for (let index = 0; index < mask.length && indexInput < input.length; index++) {
            const isReplaceValue = mask[index] === this.replacer
            const element = isReplaceValue ? input[indexInput] : mask[index]
            
            if (isReplaceValue) {
                indexInput++
            }            

            formattedString += element
        }

        return formattedString
    }

}