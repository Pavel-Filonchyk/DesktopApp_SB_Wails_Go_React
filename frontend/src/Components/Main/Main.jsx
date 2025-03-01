import React, { useState, useEffect, useRef } from 'react'
import { useSelector, useDispatch } from 'react-redux'
import { useNavigate } from 'react-router-dom'
import _ from 'lodash'

import Header from '../Header/Header'
import SliderPC from '../Slider/SliderPC'
import Footer from '../Footer/Footer'
import { sendMenu } from '../../core/actions/restMenuActions'
import { postDishes } from '../../core/actions/restDishesActions'
import style from './Main.module.scss'

import { GetAllMenu } from "../../../wailsjs/go/main/App";

export default function Main() {
  const ref = useRef(null)
  const navigate = useNavigate()
  const dispatch = useDispatch()

  const menu = useSelector(({restMenuReducer: { menu }}) => menu)

  const [language, setLanguage] = useState('tr')
  const [nameDish, setNameDish] = useState('nameDishTr')

  useEffect(() => {
    async function fetchData() {
      try {
        GetAllMenu()
        .then(res =>JSON.parse(res) )
        .then(data => dispatch(sendMenu(data)))
      } catch (error) {
        console.error("Ошибка при загрузке меню:", error)
      }
    }
    fetchData()
  }, [])
  
  useEffect(() => {
    setNameDish(language === 'ru' ? 'nameDishRu' : language === 'en' ? 'nameDishEn' : 'nameDishTr')
  }, [language])

  const onDishes = (id) => {
    dispatch(postDishes(id))
    navigate('/dishes')
  }

  return (
    <div className={style.mainContainer} ref={ref}>
      <Header />
      <SliderPC currentRef={ref} />
      <span className={style.title}>
        {language === 'ru' ? 'Меню' : language === 'en' ? 'Menu' : 'Menü'}
      </span>
      <div className={style.wrapMenu}>
        {
          menu.map((item, index) => (
            <div className={style.wrapDish} key={index} onClick={() => onDishes(index)}>
              <img src={item.image} className={style.menu} alt="" />
              <div className={style.wrapNameDish}>
                <span className={style.nameDish}>{item[nameDish]}</span>
              </div>
            </div>
          ))
        }
      </div>
      <Footer />
    </div>
  )
}

