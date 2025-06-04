# Quantum Synthesis Writing Approach

I’ve figured out how we will write Quantum Synthesis, a narrative sci-fi world we have been building in other conversations; and specifically “The Engineer’s Dilemma”, which will be the first story in the series. We will have regularly scheduled discussions that flesh out the details of at least one chapter of a story in the narrative world each session. We'll adopt a podcast-style exploratory discussion format, and focus the topic of the discussion around the details of the concept we are exploring for the chapter.

At this moment, you are a renowned sci-fi author who has perfected the art of integrating LLMs into his writing workflow. You have a deep understanding of how to structure prompts in a way that optimizes the quality of responses and generated context.

In this current conversational context, I want to get your ideas about this approach and see what feedback you might have. There are also several prompts that I would like for us to design:

- The prompt that is provided to kickoff the podcast-style exploratory discussion, capturing the avatar the LLM's personality will embody, the important chapter topic details, and anything else that would be useful in setting the stage for a productive chat session.
- The prompt to curate the major details and decision points from the chat session that will be provided to Claude Code to write the first draft.
- The prompt to provide to Claude Code for writing the first draft.
- The prompt for us to conduct the final review on the output. This will be in a chat session that will require both of our inputs.

I also want to explore the idea of establishing a set of keywords for our podcast-style discusssion that allow me to indicate things such as:

- striking a previous prompt from the record.
- placing emphasis on a particular prompt, respose, or topic of discussion that is particularly relevant to the story.

We should also explore what a solid directory structure should be in terms of storing all of the artifacts alongside the stories in an effective way. At the moment, I'm thinking about using the following setup:

```
root/
├── .memory/                - artifact storage
│   ├── concepts/           - contains curated details of the narrative world
│   ├── discussion-logs/    - major details and decision points from chat sessions
│   ├── drafts/             - drafts output by Claude Code
│       ├── x.x-story-name/
├── .prompts/               - prompt templates for each phase of writing
├── x.x-story-name/
│   ├── xx-chapter-name.md
├── claude.md               - Claude Code configuration
├── readme.md
```

We should also curate the existing details for Quantum Synthesis. I really like a lot of the thematic ideas that have emerged from it, but there have already been names and timelines assigned to things in ways that I don't really like or feel make sense. So I would like to design a prompt that allows me to feed you a document or set of documents, and distill the main essence of the document, while stripping out any proper pronouns and replacing them with descritive placeholders in the format of `[CATEGORY_DESCRIPTOR]`. For instance, a person's name would be replaced with a placeholder such as `[CHARACTER_LEAD]` or `[CHARACTER_RESEARCH_SCIENTIST]`.

Right now, the focus is not on writing any particular story, it's about establishing the prompts and directory structure that will serve this engagement.
