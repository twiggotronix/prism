import { render } from "@testing-library/react";

import { QueryClientProvider } from "@tanstack/react-query";
import { FormProvider, useForm } from "react-hook-form";
import { queryClientFactory } from "../../../../../tests/query-client-factory";
import TextInput from "./text-input";

describe("TextInput", () => {
    const queryClient = queryClientFactory();
    const renderTextInput = () => {
        const TextInputComponent = () => {
            const methods = useForm();
            return (
                <QueryClientProvider client={queryClient}>
                    <FormProvider {...methods}>
                        <TextInput
                            id="testId"
                            label="a test label"
                            name="test"
                            placeholder="a test placeholder"
                        />
                    </FormProvider>
                </QueryClientProvider>
            );
        };
        return render(<TextInputComponent />);
    };
    it("should render successfully", () => {
        const { baseElement } = renderTextInput();
        expect(baseElement).toBeTruthy();
    });
});
