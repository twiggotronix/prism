import { render } from "@testing-library/react";

import { QueryClientProvider } from "@tanstack/react-query";
import { FormProvider, useForm } from "react-hook-form";
import { queryClientFactory } from "../../../../../tests/query-client-factory";
import SelectInput from "./select-input";

describe("SelectInput", () => {
    const queryClient = queryClientFactory();
    const renderSelectInput = () => {
        const SelectInputComponent = () => {
            const methods = useForm();
            return (
                <QueryClientProvider client={queryClient}>
                    <FormProvider {...methods}>
                        <SelectInput
                            id="testId"
                            label="a test label"
                            name="test"
                            options={[{ label: "test label", value: "1" }]}
                        />
                    </FormProvider>
                </QueryClientProvider>
            );
        };
        return render(<SelectInputComponent />);
    };
    it("should render successfully", () => {
        const { baseElement } = renderSelectInput();
        expect(baseElement).toBeTruthy();
    });
});
