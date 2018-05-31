# ELAG 2018 Library Data APIs Bootcamp

[![Join the chat at https://gitter.im/elag18apis/Lobby](https://badges.gitter.im/elag18apis/Lobby.svg)](https://gitter.im/elag18apis/Lobby?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

Welcome to the open repository, documentation and materials for the ELAG 2018 Library Data APIs Bootcamp!

* When: Monday, June 4th, 2018, 10:00 - 16:00
* Where: [**National Library of Technology**](https://www.elag2018.org/how-to-access-ntk/)
* Workshop Materials: [github.com/cmh2166/elag18apis](https://github.com/cmh2166/elag18apis)
* Workshop Slides:

## About the Bootcamp

A growing number of people are seeing the need to evolve library data systems to a data-forward microservices architecture. But, how to get there? This bootcamp gives an overview of a data API & service example, going from design to development to deployment. *It’s not meant to be an in-depth dive of every topic therein, but link these domains and topics together.* We hope that participants would then leave with more comfort on how to start separating out services and data needs for an evolution to clear RESTful APIs and scalable microservices.

[free space where you can throw in any other buzzwords you'd like to see.]

This bootcamp is a registration-only workshop as part of [ELAG 2018 in Prague](https://www.elag2018.org/). If you have questions about the workshop specifically, you can [contact the workshop leaders using the information below](#contact-before-during-after-the-bootcamp).

## Bootcamp Schedule

Time               | Topic                                                          | Leader(s)
------------------ | -------------------------------------------------------------- | ------------------------------------------
**9-9:10 AM**      | Workshop Introduction, Logistics, Goals (10 minutes)           | [Christina](mailto:cmharlow@stanford.edu)
**9:10-9:25 AM**   | Spark Theory: Optimal use cases for Spark (15 minutes)         | [Audrey](mailto:audrey@dp.la)
**9:25-9:40 AM**   | Spark Theory: Spark Architecture (sparkitecture?) (15 minutes) | [Michael](mailto:michael@dp.la)
**9:40-9:55 AM**   | Spark Theory: RDD vs. DataFrame APIs (15 minutes)              | [Scott](mailto:scott@dp.la)
**9:55-10:15 AM**  | Spark Practice: Env/setup (20 minutes)                         | [Mark](mailto:mb@dp.la) & [Justin](mailto:jcoyne@stanford.edu)
**10:15-10:30 AM** | break (15 minutes)                                             | n/a
**10:30-10:50 AM** | Spark Practice: Working with spark-shell (20 minutes)          | [Christina](mailto:cmharlow@stanford.edu) & [Audrey](mailto:audrey@dp.la)
**10:50-11:10 AM** | Spark Practice: Working with zeppelin (20 minutes)             | [Christina](mailto:cmharlow@stanford.edu) & [Audrey](mailto:audrey@dp.la)
**11:10-11:45 AM** | Spark Practice: Interacting with Real World Data (20 minutes)  | Whole Group
**11:45-Noon**     | Examples & Wrap-Up (30 minutes)                                | Whole Group

## Contact Before, During, After the Bootcamp

If you have questions or concerns leading up to or after the bootcamp, please user our Gitter channel to ask or open an issue on this GitHub repository, particularly with any questions dealing with workshop preparation or any installation issues. This allows multiple workshop leaders to respond as able, and other participants can also learn.
- Open an issue [here](https://github.com/cmh2166/elag18apis/issues);
- Join our Gitter channel for chatting [here](https://gitter.im/elag18apis/Lobby?utm_source=share-link&utm_medium=link&utm_campaign=share-link).

Both require that you login or create a free account with GitHub.

During the workshop, we will indicate the best ways to get help or communicate a question/comment - however, this bootcamp is intended to be informal, so feel free to speak up or indicate you have a question at any time.

## Our Expectations of You

### Social & Mental Prep

To keep this workshop a happy, safe and inclusive space, we ask that you review and follow [the Recurse Center Social Rules (aka Hacker School Rules)](https://www.recurse.com/manual#sub-sec-social-rules).

We also ask that you come willing to take part and help out where you may have expertise to share!

### Technical Prep

We request that all participants:
- You should bring a laptop with internet connection & modern web browser;
- On that laptop, please already have:
    - have our workshop GitHub repository on your computer (with mechanism to update / pull down latest changes on Monday morning).
    - installed the latest, stable version of [Docker Community Edition](https://www.docker.com/community-edition);
    - installed latest stable Go on said laptop & set up your workspace.
    - set up a free AWS account & awscli on said laptop for said account.
    - installed localstack (requires python) on said laptop for your Go workspace.

We will use our lunch break to help folks catch up or debug these installation requirements as we are able, though the morning requires Go be set up. If you're short on time or other needs, you need at least Go, localstack, and docker for following along with local development.

If you have any issues with the above, please contact us ASAP using the [communication methods detailed above](#contact-before-during-after-the-bootcamp).
