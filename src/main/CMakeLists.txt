#rexx main module CmakeLists.txt -- rexx

set(APP_REXX rexx)

set(APP_REXX_SRC
    main.cpp
    configure.cpp
    ../api/web_server.cpp
    ../api/web_server_query.cpp
    ../api/web_server_update.cpp
    ../api/web_server_command.cpp
    ../api/web_server_helper.cpp
    ../api/websocket_server.cpp
    ../api/console.cpp
)

set(INNER_LIBS rexx_glue rexx_ledger rexx_consensus rexx_overlay rexx_common rexx_utils rexx_proto rexx_http rexx_ed25519 rexx_monitor)
set(V8_LIBS v8_base v8_libbase v8_external_snapshot v8_libplatform v8_libsampler icui18n icuuc inspector)

#Generate executable files
add_executable(${APP_REXX} ${APP_REXX_SRC})

#Specify dependent libraries for target objects
	
IF (${OS_NAME} MATCHES "OS_LINUX")  
    MESSAGE(STATUS "current platform: Linux ")  
	target_link_libraries(${APP_REXX}
    -Wl,-dn ${INNER_LIBS} -Wl,--start-group ${V8_LIBS} -Wl,--end-group ${REXX_DEPENDS_LIBS} ${REXX_LINKER_FLAGS})
ELSE ()  
	MESSAGE(STATUS "current platform: MAC ")  
	add_definitions(${REXX_LINKER_FLAGS})
	target_link_libraries(${APP_REXX} ${INNER_LIBS} ${V8_LIBS} ${REXX_DEPENDS_LIBS})
ENDIF () 


#Specify compiling options for target objets
target_compile_options(${APP_REXX}
    PUBLIC -std=c++11 
    PUBLIC -DASIO_STANDALONE
    PUBLIC -D_WEBSOCKETPP_CPP11_STL_
    PUBLIC -D${OS_NAME}
)

install(TARGETS ${APP_REXX}
    RUNTIME DESTINATION ${REXXCHAIN}/bin
)
