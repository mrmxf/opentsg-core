# tpg core

tpg-core contains all the engine and core functionality for running openTSG.

## ColourSpace Documentation

If a colour space is declared than the images use the colour space functionality. This is an *image.NRGBA64 wrapped with a
colour space, so that if a colour from a different space is added to the canvas
it can be transformed using one of the builtin transform functions.
These combine with the CNRGBA64 colours, which are the same as color.NRGBA64 but with a colourspace.
When setting a colourspace aware image with a colour space aware colour, the colours are transformed,
to match the destination (image) colourspace, on a per pixel basis.

This transformations are included in the NRGBA64 Set method and when using the Draw and DrawMask functions.

Different transformation methods will be implemented. Currently only matrix transformations are used.
to go from RGB to XYZ space and then to XYZ to RGB.


To add colour space to opentsg add the following json to the base image json.
Then add the same json code to any widget you would like to use a colourspace.

```json
"ColorSpace": {"ColorSpace" : "rec709"}
```

Current available colour spaces are :

- "rec709"
- "rec2020"
- "p3"
- "rec601"

## Factories and metadata

The input file for OpenTSG is called a factory, and can contain 1 or more references to other files.
With nesting available for the files.

When generating the widgets, the factories are processed in a depth first manner. That means every time a URI is
encountered its children and any further children are processed, before its siblings in the factory.

Each factory or widget declares which metadata keys it uses, with the "args" key
(this can be no keys).
On the generation of the widgets and factories the base metadata values
for every unique dot path are set using these keys.
This is where metadata is split from the inline update and stored in the metadata "bucket".
This base metadata "bucket" is not overwritten by later updates and is
generated on a per frame basis. It is used for applying metadata
updates to the widgets.
The workflow is the widget gets its argument keys, it searches these
keys in the metadata bucket of its parents, overwriting more generic
metadata with more specific as you proceed along the parents.
Locally declared metadata for the update will then overwirte this base metadata layer.

Wdigets can inherit any metadata that matches the declared argument keys, from their parents.
With more specific metadata overwriting previous values.

Then as the dotpath and array updates are applied, they will use these metadata values, unless
a new metadata value is called as part of that dot path.

The input factory does not have declared metadata.

## Getting started

To make it easy for you to get started with GitLab, here's a list of recommended next steps.

Already a pro? Just edit this README.md and make it your own. Want to make it easy? [Use the template at the bottom](#editing-this-readme)!

## Add your files

- [ ] [Create](https://docs.gitlab.com/ee/user/project/repository/web_editor.html#create-a-file) or [upload](https://docs.gitlab.com/ee/user/project/repository/web_editor.html#upload-a-file) files
- [ ] [Add files using the command line](https://docs.gitlab.com/ee/gitlab-basics/add-file.html#add-a-file-using-the-command-line) or push an existing Git repository with the following command:

```
cd existing_repo
git remote add origin https://github.com/mmTristan/opentsg-core.git
git branch -M main
git push -uf origin main
```

## Integrate with your tools

- [ ] [Set up project integrations](https://github.com/mmTristan/opentsg-core/-/settings/integrations)

## Collaborate with your team

- [ ] [Invite team members and collaborators](https://docs.gitlab.com/ee/user/project/members/)
- [ ] [Create a new merge request](https://docs.gitlab.com/ee/user/project/merge_requests/creating_merge_requests.html)
- [ ] [Automatically close issues from merge requests](https://docs.gitlab.com/ee/user/project/issues/managing_issues.html#closing-issues-automatically)
- [ ] [Enable merge request approvals](https://docs.gitlab.com/ee/user/project/merge_requests/approvals/)
- [ ] [Automatically merge when pipeline succeeds](https://docs.gitlab.com/ee/user/project/merge_requests/merge_when_pipeline_succeeds.html)

## Test and Deploy

Use the built-in continuous integration in GitLab.

- [ ] [Get started with GitLab CI/CD](https://docs.gitlab.com/ee/ci/quick_start/index.html)
- [ ] [Analyze your code for known vulnerabilities with Static Application Security Testing(SAST)](https://docs.gitlab.com/ee/user/application_security/sast/)
- [ ] [Deploy to Kubernetes, Amazon EC2, or Amazon ECS using Auto Deploy](https://docs.gitlab.com/ee/topics/autodevops/requirements.html)
- [ ] [Use pull-based deployments for improved Kubernetes management](https://docs.gitlab.com/ee/user/clusters/agent/)
- [ ] [Set up protected environments](https://docs.gitlab.com/ee/ci/environments/protected_environments.html)

***

# Editing this README

When you're ready to make this README your own, just edit this file and use the handy template below (or feel free to structure it however you want - this is just a starting point!).  Thank you to [makeareadme.com](https://www.makeareadme.com/) for this template.

## Suggestions for a good README
Every project is different, so consider which of these sections apply to yours. The sections used in the template are suggestions for most open source projects. Also keep in mind that while a README can be too long and detailed, too long is better than too short. If you think your README is too long, consider utilizing another form of documentation rather than cutting out information.

## Name
Choose a self-explaining name for your project.

## Description
Let people know what your project can do specifically. Provide context and add a link to any reference visitors might be unfamiliar with. A list of Features or a Background subsection can also be added here. If there are alternatives to your project, this is a good place to list differentiating factors.

## Badges
On some READMEs, you may see small images that convey metadata, such as whether or not all the tests are passing for the project. You can use Shields to add some to your README. Many services also have instructions for adding a badge.

## Visuals
Depending on what you are making, it can be a good idea to include screenshots or even a video (you'll frequently see GIFs rather than actual videos). Tools like ttygif can help, but check out Asciinema for a more sophisticated method.

## Installation
Within a particular ecosystem, there may be a common way of installing things, such as using Yarn, NuGet, or Homebrew. However, consider the possibility that whoever is reading your README is a novice and would like more guidance. Listing specific steps helps remove ambiguity and gets people to using your project as quickly as possible. If it only runs in a specific context like a particular programming language version or operating system or has dependencies that have to be installed manually, also add a Requirements subsection.

## Usage
Use examples liberally, and show the expected output if you can. It's helpful to have inline the smallest example of usage that you can demonstrate, while providing links to more sophisticated examples if they are too long to reasonably include in the README.

## Support
Tell people where they can go to for help. It can be any combination of an issue tracker, a chat room, an email address, etc.

## Roadmap
If you have ideas for releases in the future, it is a good idea to list them in the README.

## Contributing
State if you are open to contributions and what your requirements are for accepting them.

For people who want to make changes to your project, it's helpful to have some documentation on how to get started. Perhaps there is a script that they should run or some environment variables that they need to set. Make these steps explicit. These instructions could also be useful to your future self.

You can also document commands to lint the code or run tests. These steps help to ensure high code quality and reduce the likelihood that the changes inadvertently break something. Having instructions for running tests is especially helpful if it requires external setup, such as starting a Selenium server for testing in a browser.

## Authors and acknowledgment
Show your appreciation to those who have contributed to the project.

## License
For open source projects, say how it is licensed.

## Project status
If you have run out of energy or time for your project, put a note at the top of the README saying that development has slowed down or stopped completely. Someone may choose to fork your project or volunteer to step in as a maintainer or owner, allowing your project to keep going. You can also make an explicit request for maintainers.
