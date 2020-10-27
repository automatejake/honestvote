# honestvote.io

## Purpose

Elections are an important fabric of modern society. They are responsible for choosing leadership of school boards that set the direction for our communities, executives that set the direction for publicly traded companies, and politicians that set the direction of our nation. An equitable society hinges on the ability to conduct safe, transparent, and anonymous elections. These issues have always raised concern since the dawn of organized civilization. Recently, concerns have arisen that the security of US elections may have been compromised and that governments in some countries may be repressing the true results of elections. Honestvote eliminates many of these worries.

This project was created to explore alternative methods of voting and our goal is to have West Chester University use our blockchain for student elections in the Spring of 2020.

## System Outline

### Decentralization

 - How the ideal form of Honestvote looks 
	 - Blocks of transactions are approved by trusted institutions that validate 
	 - A node that requests to become a proof of authority validator ought to pay a fixed one time fee to the network that deters DDoS attacks and network spam 
	 - A node that requests to become a proof of authority validator must create a custom .txt DNS record to prove ownership of a domain 
	 - A randomly rotating number of existing validators will be selected as auditors and vote on whether or not the validator should be allowed into the system using Honestvote’s existing voting mechanism 
		 - Entropy is not predetermined because this would allow activists to predict when they were able to vote one way or another 
		 - Amount of auditors is dynamically determined by a formula. Purpose is so that not every single validator has to vote on every single new validator joining the system and also to limit the power of the majority. 
      - The current validator of any given block is determined by a round robin formula based on time 
      - Validators are the only ones that can declare elections 
         -  Nodes that are not validators do not have permission to declare elections on the Honestvote network
         - Validators must pay a network fee in order to declare an election which prevents DDoS attacks and network spam 
       - Dishonest validators are removed from the system 
          - If a block is proposed and contains dishonest information, a validator is removed as an administrator and is no longer allowed to declare elections 
         - If the transaction or block signature is incorrect or comes from an identity other than the sender, the administrator privileges are revoked 
         - If a transaction is a double spend, the administrator privileges are revoked If a transaction is for an election that has already ended, the administrator privileges are revoked
         -  If a vote does not have a corresponding registration, the administrator privileges are revoked 
         - If an administrator does not remove a dishonest validator from the system or attempts to add a validator to the system, this will result in an incorrect block causing the administrator privileges to be revoked 
         - If an administrator attempts to alter the total election history, the administrator privileges are revoked 
         - Dishonest node is no longer allowed to declare elections 
         - Dishonest node is revealed to the public 
       - Any node is able to connect to Honestvote to get the most updated information for audit purposes, they just cannot automatically participate in consensus 
       - An elected body is chosen every year that is responsible for updating the security of Honestvote algorithms and encryption 
         - Elected body is chosen by all validators in an election using Honestvote’s existing voting mechanism 
         - Two code changes are proposed per year by this body and voted upon by all validators in the Honestvote system 
  - What Honestvote already does 
      - Blocks of transactions are approved by trusted institutions that validate 
        - All validators vote on whether or not the validator should be allowed into the system using Honestvote’s existing voting mechanism 
         - The current validator of any given block is determined by a round robin formula based on time Validators are the only ones that can declare elections 
         - Nodes that are not validators do not have permission to declare elections on the Honestvote network 
         - Dishonest validators are removed from the system 
         - If a block is proposed and contains dishonest information, a validator is removed as an administrator and is no longer allowed to declare elections 
         - If the transaction or block signature is incorrect or comes from an identity other than the sender, the administrator privileges are revoked If a transaction is a double spend, the administrator privileges are revoked 
         - If a transaction is for an election that has already ended, the administrator privileges are revoked 
         - If a vote does not have a corresponding registration, the administrator privileges are revoked 
         - If an administrator does not remove a dishonest validator from the system or attempts to add a validator to the system, this will result in an incorrect block causing the administrator privileges to be revoked 
         - If an administrator attempts to alter the total election history, the administrator privileges are revoked 
         - Dishonest node is no longer allowed to declare elections 
         - Any node is able to connect to Honestvote to get the most updated information for audit purposes, they just cannot automatically participate in consensus 
       ### Anonymous 
   - How the ideal form of Honestvote will look 
      - Registrant information is private to the administrator of each election 
      - No one sees registration information besides the administrator 
      - Registration data is kept in private database of administrator and connected to administrator’s consensus node, where it stay private to that consensus node 
      - Registrant encrypts registration information with the public key of the administrator so that this information is not revealed to any party in the network besides the administrator 
      - No one can issue a registration transaction besides the administrator 
       - No one can dictate how a registration authority performs registration 
       - IP Masking Prevents a voter’s identity from being exposed based on their ip address MAC address masking 
        - Prevents a voter’s identity from being exposed to an ISP based on their MAC address SSL makes sure information sent over the internet stays private 
         - Traceable ring signatures Ring signatures ensure that the administrator, which acts as the registration authority, cannot trace back a vote to its sender based on public key 
         - Double spending is impossible because a tag prevents this from happening 
    - What Honestvote already does 
         - Registrant information is private to the administrator of each election 
         - No one sees registration information besides the administrator 
         - Registration data is kept in private database of administrator and connected to administrator’s consensus node, where it stay private to that consensus node 
         - No one can issue a registration transaction besides the administrator 
         - No one can dictate how a registration authority performs registration SSL makes sure information sent over the internet stays private 
  ### Transparent 
   - How the ideal form of Honestvote will look 
        - Any node is able to connect to the Honestvote network and get an audit log of the entire transaction history.  
        - This contains elections, registrations, and votes 
        - Voters can see the election results in real time, knowing for certainty that their vote has not been tampered with 
        - Voters can validate that their own vote has not been tampered with and was counted correctly forever 
        - No one is able to hide election data 
    - What Honestvote already does 
         - Any node is able to connect to the Honestvote network and get an audit log of the entire transaction history 
         - This contains elections, registrations, and votes 
         - Voters can see the election results in real time, knowing for certainty that their vote has not been tampered with 
         - Voters can validate that their own vote has not been tampered with and was counted correctly forever 
         - No one is able to hide election data Accessible Cross platform mobile app Mobile friendly web application Voter portal is accessible from website 
         - Very simple to navigate

## Communication

Join our Slack  [here](https://join.slack.com/t/honestvote/shared_invite/enQtNzc3MzAxNzkxMDEzLWZjNTUyZTcxNzRiNTUxYjFkNzQ0ZTJiNjFkNWUwMzdhOGE2YzllNGVhODE2NGYzNzY3ZDVhNDA3N2Q4YWRiZTg)

Participate in our Daily Scrums at 10:00 PM EST Monday-Thursday by joining  [here](https://meet.google.com/ssp-djge-nmx)

Watch us live on Twitch during our Bi-weekly Sprint Retrospective and Sprint Planning during our 12 hour programming session from 9:00 AM to 9:00 PM held every other Saturday

Check out our website at  [honestvote.io](https://honestvote.io/)


