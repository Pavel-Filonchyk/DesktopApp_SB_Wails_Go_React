import React, { useState, useEffect } from 'react'
import { useSelector } from 'react-redux'
import { useNavigate } from 'react-router-dom'

import style from './Slider.module.scss'

import img1tr from './images/PC/1.tr.png'
import img2tr from './images/PC/2.tr.png'
import img3tr from './images/PC/3.tr.png'

import img1ru from './images/PC/1.ru.png'
import img2ru from './images/PC/2.ru.png'
import img3ru from './images/PC/3.ru.png'

import img1en from './images/PC/1.en.png'
import img2en from './images/PC/2.en.png'
import img3en from './images/PC/3.en.png'

export default function SliderPC({ currentRef }) {
  const navigate = useNavigate();
  const language = useSelector(({ chooseItemsReducer: { language } }) => language)

  const [width, setWidth] = useState(false)
  const [calc, setCalc] = useState(0)
  const [margin, setMargin] = useState(0)

  useEffect(() => {
    setWidth(currentRef.current ? currentRef.current.offsetWidth : 0)
  }, [currentRef.current])

  useEffect(() => {
    if (calc === 0) {
      setTimeout(() => setCalc(calc + 1), 2500)
      setMargin(0);
    } else if (calc === 1) {
      setTimeout(() => setCalc(calc + 1), 2500)
      setMargin('-100%')
    } else if (calc === 2) {
      setTimeout(() => setCalc(0), 2500)
      setMargin('-200%')
    }
  }, [calc])

  const images = language === 'tr' 
    ? [img1tr, img2tr, img3tr] 
    : language === 'ru' 
    ? [img1ru, img2ru, img3ru] 
    : [img1en, img2en, img3en]

  return (
    <div className={style.wrapSlider} style={{ width: width }} onClick={() => navigate('/cart')}>
      <div style={{ display: 'flex', marginLeft: margin, transition: 'all 1s ease' }}>
        {images.map((src, index) => (
          <img key={index} src={src} style={{ width: width, height: '100%' }} alt="" />
        ))}
      </div>
    </div>
  )
}
