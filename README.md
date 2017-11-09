## AAML -> AAML aint a Markup Language

A simple compiler that lexically analyzes, parses and interpret the following text pattern into __key,value__ pairs. 

__Text Pattern :__

	name -> rojesh
	age -> 32
	friends -> robus,rajan,raman


### Grammer : 

    statementList	    :   statement
                            |   statementList

	statement	  	    :   assignmentStatement
							|   empty

	assignmentStatement :   factor ARROW expr

	expr                :   factor (COMMA factor)*
						    |   empty

	factor 				:   STRING
