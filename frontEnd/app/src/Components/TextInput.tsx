import React from 'react';

interface Props {
    name: string;
    label: string;
    value: string;
    onChange: (newValue: string) => void;
    required?: boolean;
}

const TextInput = (props: Props) => {
    const {name, label, value, onChange, required} = props;

    return (
        <label>
            <span className="label-text">{label}</span>
            <input
                type="text"
                name={name}
                value={value}
                onChange={(event) => onChange(event.target.value)}
                required={required}
            />
        </label>
    );
}

export default TextInput;
