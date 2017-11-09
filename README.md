## AAML -> AAML aint a Markup Language

A simple compiler that lexically analyzes, parses and interpret the following text pattern into __key,value__ pairs. 

### Grammer : 

    statementList	    :   statement
                            |   statementList

	statement	  	    :   assignmentStatement
							|   empty

	assignmentStatement :   factor ARROW expr

	expr                :   factor (COMMA factor)*
						    |   empty

	factor 				:   STRING


__Text Pattern :__

	name -> rojesh
	age -> 32
	friends -> robus,rajan,raman

__Output :__
    
    map[name:[rojesh] age:[32] friends:[robus rajan raman]]
    