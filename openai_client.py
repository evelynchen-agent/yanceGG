from openai import OpenAI


def main() -> None:
    client = OpenAI(api_key="xxxx")

    response = client.responses.create(
        model="gpt-4.1-mini",
        input="Say hello from the OpenAI API example.",
    )

    print(response.output_text)


if __name__ == "__main__":
    main()
