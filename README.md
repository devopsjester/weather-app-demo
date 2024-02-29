# Preface
Unlike many workshops, this one doesn't offer a "do it by the numbers" solution. Copilot suggestions are *not* copied & pasted from a data store; rather, they are created and tailored for your context, based on machine learning from many open source codebases. Copilot requires you to be the main pilot. Copilot merely *helps you* by suggesting code by your prompts.

Furthermore, GitHub Copilot works with many different languages, with better suggestions for the more popular languages and frameworks. It is recommended you pick from among Python, Go, TypeScript, C#, and Java. You may try any other language as well.

As such, your mileage may vary, and you will be expected to take the lead, with minimal guidance. Working experience as a software developer is required.

Good luck!

# Getting started
The first thing you will want to do is create a new repo. Create a repo. Name the repo something appropriate and unique, such as `<your username>-weather`.

Add a `.gitignore` appropriate to your chosen language. You may add a readme file for your convenience. 

# Codespaces
Once you've created a repo, you may want to create a Codespace within which to work. To do so, first, you will want to create a configuration file. 

Add a new file in GitHub/.devcontainer/devcontainer.json.

Below, is an example devcontainer file for Python:
```
{
    "name": "Python Project",
    "image": "python:3.11",
    "extensions": [ 
        "ms-python.python", 
        "github.copilot", 
        "github.copilot-chat" 
    ]
}
```
Commit the file to the main branch.

Next, create a new Codespace; click on the green `Code` button, and click on the `Codespaces` tab. Click `Create Codespace on main`. 

This will open visual studio code.

# Copilot Primer
## Getting started
If you set up the `devcontainer.json` file properly, the *Copilot* and *Copilot-Chat* extensions will be added automatically (see the toolbar on the left). You will need to log into Copilot with your GitHub username.

## Combine Workflows
Make the most out of Copilot by mixing up different workflows:
- Write code faster with Copilot suggestions in the editor.
- Rewrite blocks of code with the interactive editor session.
- Ask questions and iterate on bigger problems in the chat.

## Tips
Use the Control-Shift-P (Command-Shift-P on a Mac) keyboard shortcut in VSCode and type 'copilot' to find the commands you can run in the editor. Commands with keyboard shortcuts will have those marked to the right of the individual commands.

Before you get started, open the sidebar chat and ask  Copilot for "help on Copilot", to get started. 

# Prompts
Using the Copilot Chat interface, enter the following prompt:
```
Write an app that takes the zipcode as an argument, and gets the city and state from it, using zipopotam.us.
```
> zipopotam.us is a free public API that can return the city and state based on a zipcode.

Call the function you created using a zipcode passed in from the command line argument.

Example prompt:
```
Get the zipcode from the command line argument.
```
> Use the inline chat to add this under the function you created before.

Get the weather for the city and state you found.
Use the sidebar chat interface to create a function that will get the temperature, in degrees Fahrenheit, for the city and state.
Use a prompt such as:
```
Get the temperature for the city and state, using the openweathermap.org API in imperial units.
```
> The openweathermap.org API is also public and free. You will, however, have to register for it, and create an API key to use for the application. Make sure to replace any variable or constant with your actual API key.

Highlight the main code that calls the function that gets the city and state, which you've created above (or create it now, if it doesn't exist yet), and use Copilot to also add a call to the function that gets the temperature.

An example prompt might look something like:
```
Change the code to get the city and state from the zip code, and then get the temperature in Farenheit. Print the temperature in the city.
```

If everything worked so far, you should be able to call your app with a command like: `python weather-app 10001`, and get a result similar to `The temperature in New York City is 74.32 degrees Farenheit.`

# More things you can try.
In the sidebar, you can ask Copilot to `/explain` a block of code or the entire page to you. If you have errors, you can highlight a code block, and ask Copilot to `/fix` it for you. Consider mentioning a specific error message to help it figure out the context.
You can even use `/tests` to add unit tests for a selected block of code.

## Further tasks
- Try to add support for both Celsius and Farenheit units of measurement. Try creating a default behavior. Create a clear and explicit prompt to help Copilot understand what you're trying to do.
- Use the `openweathermap.org` API to gain and display additional information, such as weather conditions.
- Try adding better error handling and logging.
- Try improving the command line API.
- Try making the code accept either a zip code, or city & state, as inputs, but not both.

## Even further tasks
- in a completely new repository, create a new application that will tell a random joke. Use an API to find these.
- Try doing something  *completely* different. Perhaps use a public (or even your company's) APIs to display some information on a web or CLI app.

# Have fun!
